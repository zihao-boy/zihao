(function(vc) {

    vc.extends({
        data: {
            editAppServiceVarInfo: {
                asId: '',
                varSpec: '',
                varValue: '',
                varName: '',
                avId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editAppServiceVar', 'openEditAppServiceVarModal', function(_param) {
                vc.copyObject(_param, $that.editAppServiceVarInfo);
                $('#editAppServiceVarModel').modal('show');
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
                    avId: ''
                };
            }
        }
    });

})(window.vc);