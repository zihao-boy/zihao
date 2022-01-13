(function(vc) {

    vc.extends({
        data: {
            editAppServiceHostsInfo: {
                asId: '',
                hostname: '',
                ip: '',
                hostsId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editAppServiceHosts', 'openEditAppServiceHostsModal', function(_param) {
                vc.copyObject(_param, $that.editAppServiceHostsInfo);
                $('#editAppServiceHostsModel').modal('show');
            });
        },
        methods: {
            editAppServiceHostsValidate() {
                return vc.validate.validate({
                    editAppServiceHostsInfo: vc.component.editAppServiceHostsInfo
                }, {
                    'editAppServiceHostsInfo.hostname': [{
                        limit: "required",
                        param: "",
                        errInfo: "域名不能为空"
                    }],
                    'editAppServiceHostsInfo.ip': [{
                        limit: "required",
                        param: "",
                        errInfo: "ip不能为空"
                    }]
                });
            },
            updateAppServiceHostsInfo: function() {
                if (!vc.component.editAppServiceHostsValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appService/updateAppServiceHosts',
                    JSON.stringify(vc.component.editAppServiceHostsInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppServiceHostsModel').modal('hide');
                            vc.emit('appServiceControlHosts', 'switch', $that.editAppServiceHostsInfo);
                            vc.component.cleareditAppServiceHostsInfo();

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            cleareditAppServiceHostsInfo: function() {
                vc.component.editAppServiceHostsInfo = {
                    asId: '',
                    hostname: '',
                    ip: '',
                    hostsId: ''
                };
            }
        }
    });

})(window.vc);