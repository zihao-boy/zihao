(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addInnerNetUserInfo: {
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
            vc.on('addInnerNetUser', 'openAddInnerNetUserModal', function() {
                $('#addInnerNetUserModel').modal('show');
            });
        },
        methods: {
            addInnerNetUserValidate() {
                return vc.validate.validate({
                    addInnerNetUserInfo: vc.component.addInnerNetUserInfo
                }, {
                    'addInnerNetUserInfo.username': [{
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
                    'addInnerNetUserInfo.password': [{
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
                    'addInnerNetUserInfo.tel': [{
                            limit: "required",
                            param: "",
                            errInfo: "手机号不能为空"
                        }
                    ],
                });
            },
            saveInnerNetUserInfo: function() {
                if (!vc.component.addInnerNetUserValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                if($that.addInnerNetUserInfo.ipType == 'dhcp'){
                    $that.addInnerNetUserInfo.ip='0.0.0.0';
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addInnerNetUserInfo);
                    $('#addInnerNetUserModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/innerNet/saveInnerNetUser',
                    JSON.stringify(vc.component.addInnerNetUserInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addInnerNetUserModel').modal('hide');
                            vc.component.clearAddInnerNetUserInfo();
                            vc.emit('innerNetUserManage', 'listInnerNetUser', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddInnerNetUserInfo: function() {
                vc.component.addInnerNetUserInfo = {
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