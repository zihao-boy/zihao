(function(vc, vm) {

    vc.extends({
        data: {
            editDbInfo: {
                dbId: '',
                name: '',
                ip: '',
                port: '',
                username: '',
                password: '',
                dbName: '',
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editDb', 'openEditDbModal', function(_params) {
                vc.component.refreshEditDbInfo();
                $('#editDbModel').modal('show');
                vc.copyObject(_params, vc.component.editDbInfo);
            });
        },
        methods: {
            editDbValidate: function() {
                return vc.validate.validate({
                    editDbInfo: vc.component.editDbInfo
                }, {
                    'editDbInfo.name': [{
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
                    'editDbInfo.ip': [{
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
                    'editDbInfo.port': [{
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
                    'editDbInfo.username': [{
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
                    'editDbInfo.password': [{
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
                    'editDbInfo.dbName': [{
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
                    'editDbInfo.dbId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editDb: function() {
                if (!vc.component.editDbValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/resources/updateDb',
                    JSON.stringify(vc.component.editDbInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editDbModel').modal('hide');
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
            refreshEditDbInfo: function() {
                vc.component.editDbInfo = {
                    dbId: '',
                    name: '',
                    ip: '',
                    port: '',
                    username: '',
                    password: '',
                    dbName: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);