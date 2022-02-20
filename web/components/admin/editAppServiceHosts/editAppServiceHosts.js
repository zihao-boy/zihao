(function(vc) {

    vc.extends({
        data: {
            editAppServiceHostsInfo: {
                asId: '',
                hostname: '',
                ip: '',
                hostsId: '',
                isMore: '2',
                asIds: [],
                appServices: [],
                asGroupId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editAppServiceHosts', 'openEditAppServiceHostsModal', function(_param) {
                vc.copyObject(_param, $that.editAppServiceHostsInfo);
                $that.editAppServiceHostsInfo.asGroupId = _param.asGroupId;
                $that.editAppServiceHostsInfo.asIds.push(_param.asId)
                $('#editAppServiceHostsModel').modal('show');
                $that._loadEditSelectAppService();
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
                let _isMore = $that.editAppServiceHostsInfo.isMore;
                if (_isMore == 1) {
                    $that.editAppServiceHostsInfo.asId = $that.editAppServiceHostsInfo.asIds.toString();
                }

                if (!$that.editAppServiceHostsInfo.asId) {
                    vc.toast('请选择应用');
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
                    hostsId: '',
                    isMore: '2',
                    asIds: [],
                    appServices: [],
                    asGroupId: ''
                };
            },
            _loadEditSelectAppService: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 100,
                        asGroupId: $that.editAppServiceHostsInfo.asGroupId
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function(json, res) {
                        let _appVersionJobManageInfo = JSON.parse(json);
                        vc.component.editAppServiceHostsInfo.appServices = _appVersionJobManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);