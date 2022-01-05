(function(vc) {

    vc.extends({
        data: {
            fasterDeployInfo: {
                appName: '',
                deployType: '',
                packageId: '',
                shellPackageId: '',
                asGroupId: '',
                asGroups: [],
                hostGroups: [],
                groupId: '',
                asDeployType: '1001',
                hosts: [],
                hostId: '',
                openPort: ''
            }
        },
        _initMethod: function() {
            $that._listAddAppVarGroups();
            $that._listAddHostGroups();
        },
        _initEvent: function() {
            vc.on('fasterDeploy', 'fileName', function(_param) {
                $that.fasterDeployInfo.packageId = _param.id;
            })
        },
        methods: {
            fasterDeployValidate() {
                return vc.validate.validate({
                    fasterDeployInfo: vc.component.fasterDeployInfo
                }, {
                    'fasterDeployInfo.appName': [{
                            limit: "required",
                            param: "",
                            errInfo: "应用名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "应用名称太长"
                        },
                    ],
                    'fasterDeployInfo.asDeployType': [{
                            limit: "required",
                            param: "",
                            errInfo: "部署方式不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "部署方式格式错误"
                        },
                    ],
                });
            },
            saveFasterDeployInfo: function() {
                if (!vc.component.fasterDeployValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                if ($that.fasterDeployInfo.asDeployType == '1001') {
                    $that.fasterDeployInfo.asDeployId = $that.fasterDeployInfo.groupId
                } else {
                    $that.fasterDeployInfo.asDeployId = $that.fasterDeployInfo.hostId
                }

                vc.http.apiPost(
                    '/appService/saveFasterDeploy',
                    JSON.stringify(vc.component.fasterDeployInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            vc.goBack();
                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            _goBack: function() {
                vc.goBack();
            },
            _listAddAppVarGroups: function(_page, _rows) {
                var param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppVarGroup',
                    param,
                    function(json, res) {
                        var _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.fasterDeployInfo.asGroups = _appVarGroupManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _listAddHostGroups: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/host/getHostGroup',
                    param,
                    function(json, res) {
                        let _hostGroupManageInfo = JSON.parse(json);
                        vc.component.fasterDeployInfo.hostGroups = _hostGroupManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _changeHostGroup: function() {
                $that._listAddHosts();
            },
            _listAddHosts: function() {
                var param = {
                    params: {
                        page: 1,
                        row: 100,
                        groupId: $that.fasterDeployInfo.groupId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function(json, res) {
                        var _hostManageInfo = JSON.parse(json);
                        vc.component.fasterDeployInfo.hosts = _hostManageInfo.data;

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openFasterDeployLog: function() {
                vc.jumpToPage('/index.html#/pages/admin/fasterDeployLog')
            }
        }
    });

})(window.vc);