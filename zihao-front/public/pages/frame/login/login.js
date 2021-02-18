(function (vc) {
    vc.extends({
        data: {
            loginInfo: {
                logo: '',
                username: 'wuxw',
                passwd: 'admin',
            }
        },
        _initMethod: function () {
            vc.component.clearCacheData();
        },
        _initEvent: function () {
            vc.component.$on('errorInfoEvent', function (_errorInfo) {
                vc.component.loginInfo.errorInfo = _errorInfo;
                console.log('errorInfoEvent 事件被监听', _errorInfo)
            });

            vc.component.$on('validate_code_component_param_change_event', function (params) {
                for (var tmpAttr in params) {
                    vc.component.loginInfo[tmpAttr] = params[tmpAttr];
                }
                console.log('errorInfoEvent 事件被监听', params)
            });
        },
        methods: {
            clearCacheData: function () {
                vc.clearCacheData();
            },
            doLogin: function () {
                if (!vc.notNull(vc.component.loginInfo.username)) {
                    vc.toast('用户名不能为空');
                    return;
                }
                if (!vc.notNull(vc.component.loginInfo.passwd)) {
                    vc.toast('密码不能为空');
                    return;
                }
                vc.http.post(
                    'login',
                    'doLogin?version=2.0',
                    JSON.stringify(vc.component.loginInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _data = JSON.parse(json);
                        if (_data.hasOwnProperty('code') && _data.code != '0') {
                            vc.toast(_data.msg);
                            return;
                        }
                        vc.jumpToPage("/index.html")
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                        vc.component.loginInfo.errorInfo = errInfo;
                    });

            }
        },
        _destroyedMethod: function () {
            console.log("登录页面销毁调用");
        }
    });


})(window.vc);