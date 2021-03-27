(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addAppVersionJobInfo: {
                jobId: '',
                jobName: '',
                jobShell: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addAppVersionJob', 'openAddAppVersionJobModal', function () {
                $('#addAppVersionJobModel').modal('show');
            });
        },
        methods: {
            addAppVersionJobValidate() {
                return vc.validate.validate({
                    addAppVersionJobInfo: vc.component.addAppVersionJobInfo
                }, {
                    'addAppVersionJobInfo.jobName': [
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
                    'addAppVersionJobInfo.jobShell': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "脚本不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "1024",
                            errInfo: "脚本错误"
                        },
                    ],




                });
            },
            saveAppVersionJobInfo: function () {
                if (!vc.component.addAppVersionJobValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addAppVersionJobInfo);
                    $('#addAppVersionJobModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/appVersion/saveAppVersionJob',
                    JSON.stringify(vc.component.addAppVersionJobInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppVersionJobModel').modal('hide');
                            vc.component.clearAddAppVersionJobInfo();
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
            clearAddAppVersionJobInfo: function () {
                vc.component.addAppVersionJobInfo = {
                    jobName: '',
                    jobShell: '',

                };
            }
        }
    });

})(window.vc);
