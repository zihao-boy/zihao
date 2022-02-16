(function(vc) {

    vc.extends({
        data: {
            addAppServiceHostsInfo: {
                asId: '',
                hostname: '',
                ip: '',
                isMore: '2',
                asIds: [],
                appServices: [],
                asGroupId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addAppServiceHosts', 'openAddAppServiceHostsModal', function(_param) {
                $that.addAppServiceHostsInfo.asId = _param.asId;
                $that.addAppServiceHostsInfo.asGroupId = _param.asGroupId;
                $that.addAppServiceHostsInfo.asIds.push(_param.asId)
                $('#addAppServiceHostsModel').modal('show');
                $that._loadSelectAppService();
            });
        },
        methods: {
            addAppServiceHostsValidate() {
                return vc.validate.validate({
                    addAppServiceHostsInfo: vc.component.addAppServiceHostsInfo
                }, {
                    'addAppServiceHostsInfo.hostname': [{
                        limit: "required",
                        param: "",
                        errInfo: "域名不能为空"
                    }],
                    'addAppServiceHostsInfo.ip': [{
                        limit: "required",
                        param: "",
                        errInfo: "ip不能为空"
                    }]
                });
            },
            saveAppServiceHostsInfo: function() {
                if (!vc.component.addAppServiceHostsValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                let _isMore = $that.addAppServiceHostsInfo.isMore;
                if (_isMore == 1) {
                    $that.addAppServiceHostsInfo.asId = $that.addAppServiceHostsInfo.asIds.toString();
                }

                if (!$that.addAppServiceHostsInfo.asId) {
                    vc.toast('请选择应用');
                    return;
                }
                vc.http.apiPost(
                    '/appService/saveAppServiceHosts',
                    JSON.stringify(vc.component.addAppServiceHostsInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppServiceHostsModel').modal('hide');
                            vc.emit('appServiceControlHosts', 'switch', $that.addAppServiceHostsInfo);
                            vc.component.clearaddAppServiceHostsInfo();

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearaddAppServiceHostsInfo: function() {
                vc.component.addAppServiceHostsInfo = {
                    asId: '',
                    hostname: '',
                    ip: '',
                    isMore: '2',
                    asIds: [],
                    appServices: [],
                    asGroupId: ''
                };
            },
            _loadSelectAppService: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 100,
                        asGroupId: $that.addAppServiceHostsInfo.asGroupId
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function(json, res) {
                        let _appVersionJobManageInfo = JSON.parse(json);
                        vc.component.addAppServiceHostsInfo.appServices = _appVersionJobManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);