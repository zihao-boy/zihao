(function (vc, vm) {

    vc.extends({
        data: {
            editAppVarInfo: {
                avId: '',
                varName: '',
                avgId: '',
                varSpec:'',
                varType: '',
                appVarGroups:[]
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editAppVar', 'openEditAppVarModal', function (_params) {
                vc.component.refreshEditAppVarInfo();
                $('#editAppVarModel').modal('show');
                $that._listEditAppVarGroups();
                vc.copyObject(_params, vc.component.editAppVarInfo);
            });
        },
        methods: {
            editAppVarValidate: function () {
                return vc.validate.validate({
                    editAppVarInfo: vc.component.editAppVarInfo
                }, {
                    'editAppVarInfo.varName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "变量名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "变量名太长"
                        },
                    ],
                    'editAppVarInfo.avgId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "变量组不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "变量组错误"
                        },
                    ],
                    'editAppVarInfo.varSpec': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "变量编码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "变量编码错误"
                        },
                    ],
                    'editAppVarInfo.varType': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "变量类型不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "变量类型错误"
                        },
                    ],
                    'editAppVarInfo.avId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "变量编码不能为空"
                        }]

                });
            },
            editAppVar: function () {
                if (!vc.component.editAppVarValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/appService/updateAppVar',
                    JSON.stringify(vc.component.editAppVarInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppVarModel').modal('hide');
                            vc.emit('appVarManage', 'listAppVar', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditAppVarInfo: function () {
                vc.component.editAppVarInfo = {
                    avId: '',
                    varName: '',
                    avgId: '',
                    varSpec:'',
                    varType: '',
                    appVarGroups:[]
                }
            },
            _listEditAppVarGroups: function () {

                var param = {
                    params: {
                        page:1,
                        row:50
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppVarGroup',
                    param,
                    function (json, res) {
                        let _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.editAppVarInfo.appVarGroups = _appVarGroupManageInfo.data;   
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc, window.vc.component);
