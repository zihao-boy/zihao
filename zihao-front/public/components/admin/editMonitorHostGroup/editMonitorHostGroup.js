(function (vc, vm) {

    vc.extends({
        data: {
            editMonitorHostGroupInfo: {
                mhgId: '',
                name: '',
                monCron: '',
                noticeType: '',
                remark: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editMonitorHostGroup', 'openEditMonitorHostGroupModal', function (_params) {
                vc.component.refreshEditMonitorHostGroupInfo();
                $('#editMonitorHostGroupModel').modal('show');
                vc.copyObject(_params, vc.component.editMonitorHostGroupInfo);
            });
        },
        methods: {
            editMonitorHostGroupValidate: function () {
                return vc.validate.validate({
                    editMonitorHostGroupInfo: vc.component.editMonitorHostGroupInfo
                }, {
                    'editMonitorHostGroupInfo.name': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "名称太长"
                        },
                    ],
                    'editMonitorHostGroupInfo.monCron': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "监控频率不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "监控频率太长"
                        },
                    ],
                    'editMonitorHostGroupInfo.noticeType': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "告警类型不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "告警类型错误"
                        },
                    ],
                    'editMonitorHostGroupInfo.remark': [
                        {
                            limit: "maxLength",
                            param: "1000",
                            errInfo: "备注太长"
                        },
                    ],
                    'editMonitorHostGroupInfo.mhgId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "组ID不能为空"
                        }]

                });
            },
            editMonitorHostGroup: function () {
                if (!vc.component.editMonitorHostGroupValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/monitor/updateMonitorHostGroup',
                    JSON.stringify(vc.component.editMonitorHostGroupInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editMonitorHostGroupModel').modal('hide');
                            vc.emit('monitorHostGroupManage', 'listMonitorHostGroup', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditMonitorHostGroupInfo: function () {
                vc.component.editMonitorHostGroupInfo = {
                    mhgId: '',
                    name: '',
                    monCron: '',
                    noticeType: '',
                    remark: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
