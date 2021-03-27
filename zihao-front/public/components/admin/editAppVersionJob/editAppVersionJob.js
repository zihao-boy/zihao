(function (vc, vm) {

    vc.extends({
        data: {
            editAppVersionJobInfo: {
                jobId: '',
                jobName: '',
                jobShell: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editAppVersionJob', 'openEditAppVersionJobModal', function (_params) {
                vc.component.refreshEditAppVersionJobInfo();
                $('#editAppVersionJobModel').modal('show');
                vc.copyObject(_params, vc.component.editAppVersionJobInfo);
            });
        },
        methods: {
            editAppVersionJobValidate: function () {
                return vc.validate.validate({
                    editAppVersionJobInfo: vc.component.editAppVersionJobInfo
                }, {
                    'editAppVersionJobInfo.jobName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "名称太长"
                        },
                    ],
                    'editAppVersionJobInfo.jobShell': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "脚本不能为空"
                        }
                    ],
                    'editAppVersionJobInfo.jobId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "构建ID不能为空"
                        }]

                });
            },
            editAppVersionJob: function () {
                if (!vc.component.editAppVersionJobValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/appVersion/updateAppVersionJob',
                    JSON.stringify(vc.component.editAppVersionJobInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppVersionJobModel').modal('hide');
                            vc.emit('appVersionJobManage', 'listAppVersionJob', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditAppVersionJobInfo: function () {
                vc.component.editAppVersionJobInfo = {
                    jobId: '',
                    jobName: '',
                    jobShell: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
