(function (vc, vm) {

    vc.extends({
        data: {
            editInnerNetUserInfo: {
                userId: '',
                username: '',
                password: '',
                tel: '',
                ip: '',
                ipType: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editInnerNetUser', 'openEditInnerNetUserModal', function (_params) {
                vc.component.refreshEditInnerNetUserInfo();
                $('#editInnerNetUserModel').modal('show');
                vc.copyObject(_params, vc.component.editInnerNetUserInfo);
                if(_params.ip =='0.0.0.0'){
                    $that.editInnerNetUserInfo.ipType = 'dhcp';
                }else{
                    $that.editInnerNetUserInfo.ipType = 'statics';
                }
            });
        },
        methods: {
            editInnerNetUserValidate: function () {
                return vc.validate.validate({
                    editInnerNetUserInfo: vc.component.editInnerNetUserInfo
                }, {
                    'editInnerNetUserInfo.username': [{
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
                    'editInnerNetUserInfo.password': [{
                        limit: "required",
                        param: "",
                        errInfo: "密码不能为空"
                    },
                    {
                        limit: "maxLength",
                        param: "64",
                        errInfo: "密码不能超过64"
                    },
                    ],
                    'editInnerNetUserInfo.tel': [{
                        limit: "required",
                        param: "",
                        errInfo: "手机号不能为空"
                    }
                    ],
                    'editInnerNetUserInfo.userId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editInnerNetUser: function () {
                if (!vc.component.editInnerNetUserValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/innerNet/updateInnerNetUser',
                    JSON.stringify(vc.component.editInnerNetUserInfo), {
                    emulateJSON: true
                },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editInnerNetUserModel').modal('hide');
                            vc.emit('innerNetUserManage', 'listInnerNetUser', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditInnerNetUserInfo: function () {
                vc.component.editInnerNetUserInfo = {
                    userId: '',
                    username: '',
                    password: '',
                    tel: '',
                    ip: '',
                    ipType: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);