(function (vc) {

    vc.extends({
        data: {
            addAppServiceHostsInfo: {
                asId: '',
                hostname: '',
                ip: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addAppServiceHosts', 'openAddAppServiceHostsModal', function (_param) {
                $that.addAppServiceHostsInfo.asId = _param.asId;
                $('#addAppServiceHostsModel').modal('show');
            });
        },
        methods: {
            addAppServiceHostsValidate() {
                return vc.validate.validate({
                    addAppServiceHostsInfo: vc.component.addAppServiceHostsInfo
                }, {
                    'addAppServiceHostsInfo.hostname': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "域名不能为空"
                        }
                    ],
                    'addAppServiceHostsInfo.ip': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "ip不能为空"
                        }
                    ]
                });
            },
            saveAppServiceHostsInfo: function () {
                if (!vc.component.addAppServiceHostsValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appService/saveAppServiceHosts',
                    JSON.stringify(vc.component.addAppServiceHostsInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppServiceHostsModel').modal('hide');
                            vc.component.clearaddAppServiceHostsInfo();
                            vc.emit('appServiceControlHosts', 'switch', $that.addAppServiceHostsInfo);
                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearaddAppServiceHostsInfo: function () {
                vc.component.addAppServiceHostsInfo = {
                    asId: '',
                    hostname: '',
                    ip: ''
                };
            }
        }
    });

})(window.vc);
