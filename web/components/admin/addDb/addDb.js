(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addDbInfo: {
                dbId: '',
                name: '',
                ip: '',
                port: '3306',
                username: '',
                password: '',
                dbName: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addDb', 'openAddDbModal', function() {
                $('#addDbModel').modal('show');
            });
        },
        methods: {
            addDbValidate() {
                return vc.validate.validate({
                    addDbInfo: vc.component.addDbInfo
                }, {
                    'addDbInfo.name': [{
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "名称不能超过64"
                        },
                    ],
                    'addDbInfo.ip': [{
                            limit: "required",
                            param: "",
                            errInfo: "ip不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "ip不能超过128"
                        },
                    ],
                    'addDbInfo.port': [{
                            limit: "required",
                            param: "",
                            errInfo: "端口不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "12",
                            errInfo: "端口不能超过12"
                        },
                    ],
                    'addDbInfo.username': [{
                            limit: "required",
                            param: "",
                            errInfo: "用户名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "用户名不能超过64"
                        },
                    ],
                    'addDbInfo.password': [{
                            limit: "required",
                            param: "",
                            errInfo: "密码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "密码不能超过128"
                        },
                    ],
                    'addDbInfo.dbName': [{
                            limit: "required",
                            param: "",
                            errInfo: "数据库不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "数据库不能超过64"
                        },
                    ],




                });
            },
            saveDbInfo: function() {
                if (!vc.component.addDbValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addDbInfo);
                    $('#addDbModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/resources/saveDb',
                    JSON.stringify(vc.component.addDbInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addDbModel').modal('hide');
                            vc.component.clearAddDbInfo();
                            vc.emit('dbManage', 'listDb', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddDbInfo: function() {
                vc.component.addDbInfo = {
                    name: '',
                    ip: '',
                    port: '3306',
                    username: '',
                    password: '',
                    dbName: '',

                };
            }
        }
    });

})(window.vc);