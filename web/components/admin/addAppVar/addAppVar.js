(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addAppVarInfo: {
                avId: '',
                varSpec: '',
                varName: '',
                avgId: '',
                varType: '',
                appVarGroups:[]
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addAppVar', 'openAddAppVarModal', function () {
                $that._listAddAppVarGroups();
                $('#addAppVarModel').modal('show');
            });
        },
        methods: {
            addAppVarValidate() {
                return vc.validate.validate({
                    addAppVarInfo: vc.component.addAppVarInfo
                }, {
                    'addAppVarInfo.varName': [
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
                    'addAppVarInfo.avgId': [
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
                    ],'addAppVarInfo.varSpec': [
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
                    'addAppVarInfo.varType': [
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




                });
            },
            saveAppVarInfo: function () {
                if (!vc.component.addAppVarValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addAppVarInfo);
                    $('#addAppVarModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/appService/saveAppVar',
                    JSON.stringify(vc.component.addAppVarInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppVarModel').modal('hide');
                            vc.component.clearAddAppVarInfo();
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
            clearAddAppVarInfo: function () {
                vc.component.addAppVarInfo = {
                    varName: '',
                    avgId: '',
                    varSpec: '',
                    varType: '',
                    appVarGroups:[]
                };
            },
            _listAddAppVarGroups: function () {

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
                        vc.component.addAppVarInfo.appVarGroups = _appVarGroupManageInfo.data;   
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);
