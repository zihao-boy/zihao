(function(vc) {

    vc.extends({
        data: {
            editAppServiceVarInfo: {
                asId: '',
                varSpec: '',
                varValue: '',
                varName: '',
                avId: '',
                isMore: '2',
                asIds: [],
                appServices: [],
                asGroupId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editAppServiceVar', 'openEditAppServiceVarModal', function(_param) {
                $that.editAppServiceVarInfo.asId = _param.asId;
                $that.editAppServiceVarInfo.asGroupId = _param.asGroupId;
                $that.editAppServiceVarInfo.asIds.push(_param.asId)
                vc.copyObject(_param, $that.editAppServiceVarInfo);
                $('#editAppServiceVarModel').modal('show');
                $that._loadEditVarSelectAppService();
            });
        },
        methods: {
            editAppServiceVarValidate() {
                return vc.validate.validate({
                    editAppServiceVarInfo: vc.component.editAppServiceVarInfo
                }, {
                    'editAppServiceVarInfo.varSpec': [{
                        limit: "required",
                        param: "",
                        errInfo: "变量名不能为空"
                    }],
                    'editAppServiceVarInfo.varValue': [{
                        limit: "required",
                        param: "",
                        errInfo: "变量值不能为空"
                    }],
                    'editAppServiceVarInfo.varName': [{
                        limit: "required",
                        param: "",
                        errInfo: "说明不能为空"
                    }]
                });
            },
            updateAppServiceVarInfo: function() {
                if (!vc.component.editAppServiceVarValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                let _isMore = $that.editAppServiceVarInfo.isMore;
                if (_isMore == 1) {
                    $that.editAppServiceVarInfo.asId = $that.editAppServiceVarInfo.asIds.toString();
                }

                if (!$that.editAppServiceVarInfo.asId) {
                    vc.toast('请选择应用');
                    return;
                }
                vc.http.apiPost(
                    '/appService/updateAppServiceVar',
                    JSON.stringify(vc.component.editAppServiceVarInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppServiceVarModel').modal('hide');
                            vc.emit('appServiceControlVar', 'switch', $that.editAppServiceVarInfo);
                            vc.component.cleareditAppServiceVarInfo();

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            cleareditAppServiceVarInfo: function() {
                vc.component.editAppServiceVarInfo = {
                    asId: '',
                    varSpec: '',
                    varValue: '',
                    varName: '',
                    avId: '',
                    isMore: '2',
                    asIds: [],
                    appServices: [],
                    asGroupId: ''
                };
            },
            _loadEditVarSelectAppService: function() {
                let param = {
                    params: {
                        page: 1,
                        row: 100,
                        asGroupId: $that.editAppServiceVarInfo.asGroupId
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function(json, res) {
                        let _appVersionJobManageInfo = JSON.parse(json);
                        vc.component.editAppServiceVarInfo.appServices = _appVersionJobManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);