(function(vc) {

    vc.extends({
        data: {
            editDbLinkInfo: {
                id:'',
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
            vc.on('editDbLink', 'openEditDbLinkModal', function(_param) {
                vc.copyObject(_param,$that.editDbLinkInfo);
                $('#editDbLinkModel').modal('show');
            });
        },
        methods: {
            editDbLinkValidate() {
                return vc.validate.validate({
                    editDbLinkInfo: vc.component.editDbLinkInfo
                }, {
                    'editDbLinkInfo.name': [{
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
                    'editDbLinkInfo.ip': [{
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
                    'editDbLinkInfo.dbName': [{
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
                    'editDbLinkInfo.password': [{
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
            _updateDbLink: function() {
                if (!vc.component.editDbLinkValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/dbClient/updateDbLink',
                    JSON.stringify(vc.component.editDbLinkInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editDbLinkModel').modal('hide');
                            vc.component.clearEditDbLinkInfo();
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
            clearEditDbLinkInfo: function() {
                vc.component.editDbLinkInfo = {
                    id:'',
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