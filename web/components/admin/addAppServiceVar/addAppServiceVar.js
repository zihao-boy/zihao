(function(vc) {

    vc.extends({
        data: {
            addAppServiceVarInfo: {
                asId: '',
                varSpec: '',
                varValue: '',
                varName: '',
                isMore: '2',
                asIds: [],
                appServices: [],
                asGroupId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addAppServiceVar', 'openAddAppServiceVarModal', function(_param) {
                $that.addAppServiceVarInfo.asId = _param.asId;
                $that.addAppServiceVarInfo.asGroupId = _param.asGroupId;
                $that.addAppServiceVarInfo.asIds.push(_param.asId)
                $('#addAppServiceVarModel').modal('show');
                $that._loadAddVarSelectAppService();
            });
        },
        methods: {
            addAppServiceVarValidate() {
                return vc.validate.validate({
                    addAppServiceVarInfo: vc.component.addAppServiceVarInfo
                }, {
                    'addAppServiceVarInfo.varSpec': [{
                        limit: "required",
                        param: "",
                        errInfo: "变量名不能为空"
                    }],
                    'addAppServiceVarInfo.varValue': [{
                        limit: "required",
                        param: "",
                        errInfo: "变量值不能为空"
                    }],
                    'addAppServiceVarInfo.varName': [{
                        limit: "required",
                        param: "",
                        errInfo: "说明不能为空"
                    }]
                });
            },
            saveAppServiceVarInfo: function() {
                if (!vc.component.addAppServiceVarValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                let _isMore = $that.addAppServiceVarInfo.isMore;
                if (_isMore == 1) {
                    $that.addAppServiceVarInfo.asId = $that.addAppServiceVarInfo.asIds.toString();
                }

                if (!$that.addAppServiceVarInfo.asId) {
                    vc.toast('请选择应用');
                    return;
                }

                vc.http.apiPost(
                    '/appService/saveAppServiceVar',
                    JSON.stringify(vc.component.addAppServiceVarInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppServiceVarModel').modal('hide');
                            vc.emit('appServiceControlVar', 'switch', $that.addAppServiceVarInfo);
                            vc.component.clearaddAppServiceVarInfo();

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearaddAppServiceVarInfo: function() {
                vc.component.addAppServiceVarInfo = {
                    asId: '',
                    varSpec: '',
                    varValue: '',
                    varName: '',
                    isMore: '2',
                    asIds: [],
                    appServices: [],
                    asGroupId: ''
                };
            },
            _loadAddVarSelectAppService: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 100,
                        asGroupId: $that.addAppServiceVarInfo.asGroupId
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function(json, res) {
                        let _appVersionJobManageInfo = JSON.parse(json);
                        vc.component.addAppServiceVarInfo.appServices = _appVersionJobManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);