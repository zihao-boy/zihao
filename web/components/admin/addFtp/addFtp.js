(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addFtpInfo: {
                name: '',
                ip: '',
                port: '',
                username: '',
                passwd: '',
                path: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addFtp', 'openAddFtpModal', function() {
                $('#addFtpModel').modal('show');
            });
        },
        methods: {
            addFtpValidate() {
                return vc.validate.validate({
                    addFtpInfo: vc.component.addFtpInfo
                }, {
                    'addFtpInfo.name': [{
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
                    'addFtpInfo.ip': [{
                            limit: "required",
                            param: "",
                            errInfo: "IP不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "IP不能超过128"
                        },
                    ],
                    'addFtpInfo.port': [{
                            limit: "required",
                            param: "",
                            errInfo: "端口不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "端口不能超过64"
                        },
                    ],
                    'addFtpInfo.username': [{
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
                    'addFtpInfo.passwd': [{
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
                    'addFtpInfo.path': [{
                            limit: "required",
                            param: "",
                            errInfo: "路径'不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "路径'不能超过128"
                        },
                    ],




                });
            },
            saveFtpInfo: function() {
                if (!vc.component.addFtpValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addFtpInfo);
                    $('#addFtpModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/resources/saveFtp',
                    JSON.stringify(vc.component.addFtpInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addFtpModel').modal('hide');
                            vc.component.clearAddFtpInfo();
                            vc.emit('ftpManage', 'listFtp', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddFtpInfo: function() {
                vc.component.addFtpInfo = {
                    name: '',
                    ip: '',
                    port: '',
                    username: '',
                    passwd: '',
                    path: '',

                };
            }
        }
    });

})(window.vc);