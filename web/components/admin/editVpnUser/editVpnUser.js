(function (vc, vm) {

    vc.extends({
        data: {
            editVpnUserInfo: {
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
            vc.on('editVpnUser', 'openEditVpnUserModal', function (_params) {
                vc.component.refreshEditVpnUserInfo();
                $('#editVpnUserModel').modal('show');
                vc.copyObject(_params, vc.component.editVpnUserInfo);
                if(_params.ip =='0.0.0.0'){
                    $that.editVpnUserInfo.ipType = 'dhcp';
                }else{
                    $that.editVpnUserInfo.ipType = 'statics';
                }
            });
        },
        methods: {
            editVpnUserValidate: function () {
                return vc.validate.validate({
                    editVpnUserInfo: vc.component.editVpnUserInfo
                }, {
                    'editVpnUserInfo.username': [{
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
                    'editVpnUserInfo.password': [{
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
                    'editVpnUserInfo.tel': [{
                        limit: "required",
                        param: "",
                        errInfo: "手机号不能为空"
                    }
                    ],
                    'editVpnUserInfo.userId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editVpnUser: function () {
                if (!vc.component.editVpnUserValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/vpn/updateVpnUser',
                    JSON.stringify(vc.component.editVpnUserInfo), {
                    emulateJSON: true
                },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editVpnUserModel').modal('hide');
                            vc.emit('vpnUserManage', 'listVpnUser', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditVpnUserInfo: function () {
                vc.component.editVpnUserInfo = {
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