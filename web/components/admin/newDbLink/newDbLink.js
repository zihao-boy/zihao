(function(vc) {

    vc.extends({
        data: {
            newDbLinkInfo: {
                name: '',
                ip: '',
                port: '3306',
                username: 'root',
                password: '',
                dbName: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('newDbLink', 'openNewDbLinkModal', function() {
                $('#newDbLinkModel').modal('show');
            });
        },
        methods: {
            newDbLinkValidate() {
                return vc.validate.validate({
                    newDbLinkInfo: vc.component.newDbLinkInfo
                }, {
                    'newDbLinkInfo.name': [{
                            limit: "required",
                            param: "",
                            errInfo: "连接名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "连接名格式错误"
                        },
                    ],
                    'newDbLinkInfo.ip': [{
                            limit: "required",
                            param: "",
                            errInfo: "主机名或IP不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "主机名或IP太长"
                        },
                    ],
                    'newDbLinkInfo.dbName': [{
                            limit: "required",
                            param: "",
                            errInfo: "数据库不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "数据库太长"
                        },
                    ],
                    'newDbLinkInfo.password': [{
                            limit: "required",
                            param: "",
                            errInfo: "密码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "密码太长"
                        },
                    ],
                });
            },
            _saveDbLink: function() {
                if (!vc.component.newDbLinkValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/dbClient/saveDbLink',
                    JSON.stringify(vc.component.newDbLinkInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#newDbLinkModel').modal('hide');
                            vc.component.clearNewDbLinkInfo();
                            vc.emit('mysqlClient', 'load', {});

                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            clearNewDbLinkInfo: function() {
                vc.component.newDbLinkInfo = {
                    name: '',
                    ip: '',
                    port: '3306',
                    username: 'root',
                    password: '',
                    dbName: ''
                };
            }
        }
    });

})(window.vc);