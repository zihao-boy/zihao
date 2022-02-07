(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addAppPublisherInfo: {
                publisherId: '',
                username: '',
                email: '',
                phone: '',
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addAppPublisher', 'openAddAppPublisherModal', function() {
                $('#addAppPublisherModel').modal('show');
            });
        },
        methods: {
            addAppPublisherValidate() {
                return vc.validate.validate({
                    addAppPublisherInfo: vc.component.addAppPublisherInfo
                }, {
                    'addAppPublisherInfo.username': [{
                            limit: "required",
                            param: "",
                            errInfo: "发布者不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "256",
                            errInfo: "发布者不能超过256"
                        },
                    ],
                    'addAppPublisherInfo.email': [{
                            limit: "required",
                            param: "",
                            errInfo: "邮箱不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "邮箱不能超过64"
                        },
                    ],
                    'addAppPublisherInfo.phone': [{
                            limit: "required",
                            param: "",
                            errInfo: "电话不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "11",
                            errInfo: "电话不能超过11"
                        },
                    ],
                });
            },
            saveAppPublisherInfo: function() {
                if (!vc.component.addAppPublisherValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addAppPublisherInfo);
                    $('#addAppPublisherModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/soft/saveAppPublisher',
                    JSON.stringify(vc.component.addAppPublisherInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppPublisherModel').modal('hide');
                            vc.component.clearAddAppPublisherInfo();
                            vc.emit('appPublisherManage', 'listAppPublisher', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddAppPublisherInfo: function() {
                vc.component.addAppPublisherInfo = {
                    username: '',
                    email: '',
                    phone: '',

                };
            }
        }
    });

})(window.vc);