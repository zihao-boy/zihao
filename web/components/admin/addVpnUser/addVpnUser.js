(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addVpnUserInfo: {
                username: '',
                    password: '',
                    tel: '',
                    ip: '',
                    ipType:'',
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addVpnUser', 'openAddVpnUserModal', function() {
                $('#addVpnUserModel').modal('show');
            });
        },
        methods: {
            addVpnUserValidate() {
                return vc.validate.validate({
                    addVpnUserInfo: vc.component.addVpnUserInfo
                }, {
                    'addVpnUserInfo.username': [{
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
                    'addVpnUserInfo.password': [{
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
                    'addVpnUserInfo.tel': [{
                            limit: "required",
                            param: "",
                            errInfo: "手机号不能为空"
                        }
                    ],
                });
            },
            saveVpnUserInfo: function() {
                if (!vc.component.addVpnUserValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                if($that.addVpnUserInfo.ipType == 'dhcp'){
                    $that.addVpnUserInfo.ip='0.0.0.0';
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addVpnUserInfo);
                    $('#addVpnUserModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/vpn/saveVpnUser',
                    JSON.stringify(vc.component.addVpnUserInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addVpnUserModel').modal('hide');
                            vc.component.clearAddVpnUserInfo();
                            vc.emit('vpnUserManage', 'listVpnUser', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddVpnUserInfo: function() {
                vc.component.addVpnUserInfo = {
                    username: '',
                    password: '',
                    tel: '',
                    ip: '',
                    ipType:'',

                };
            },
        }
    });

})(window.vc);