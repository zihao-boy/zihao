(function (vc) {

    vc.extends({
        data: {
            addAppServiceVarInfo: {
                asId: '',
                varSpec: '',
                varValue: '',
                varName:''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addAppServiceVar', 'openAddAppServiceVarModal', function (_param) {
                $that.addAppServiceVarInfo.asId = _param.asId;
                $('#addAppServiceVarModel').modal('show');
            });
        },
        methods: {
            addAppServiceVarValidate() {
                return vc.validate.validate({
                    addAppServiceVarInfo: vc.component.addAppServiceVarInfo
                }, {
                    'addAppServiceVarInfo.varSpec': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "变量名不能为空"
                        }
                    ],
                    'addAppServiceVarInfo.varValue': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "变量值不能为空"
                        }
                    ],
                    'addAppServiceVarInfo.varName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "说明不能为空"
                        }
                    ]
                });
            },
            saveAppServiceVarInfo: function () {
                if (!vc.component.addAppServiceVarValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appService/saveAppServiceVar',
                    JSON.stringify(vc.component.addAppServiceVarInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppServiceVarModel').modal('hide');
                            vc.component.clearaddAppServiceVarInfo();
                            vc.emit('appServiceControlVar', 'switch', $that.addAppServiceVarInfo);
                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearaddAppServiceVarInfo: function () {
                vc.component.addAppServiceVarInfo = {
                    asId: '',
                    varSpec: '',
                    varValue: '',
                    varName:''
                };
            }
        }
    });

})(window.vc);
