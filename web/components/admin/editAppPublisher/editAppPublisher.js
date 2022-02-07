(function(vc, vm) {

    vc.extends({
        data: {
            editAppPublisherInfo: {
                publisherId: '',
                username: '',
                email: '',
                phone: '',
                extPublisherId:''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editAppPublisher', 'openEditAppPublisherModal', function(_params) {
                vc.component.refreshEditAppPublisherInfo();
                $('#editAppPublisherModel').modal('show');
                vc.copyObject(_params, vc.component.editAppPublisherInfo);
            });
        },
        methods: {
            editAppPublisherValidate: function() {
                return vc.validate.validate({
                    editAppPublisherInfo: vc.component.editAppPublisherInfo
                }, {
                    'editAppPublisherInfo.username': [{
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
                    'editAppPublisherInfo.email': [{
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
                    'editAppPublisherInfo.phone': [{
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
                    'editAppPublisherInfo.publisherId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editAppPublisher: function() {
                if (!vc.component.editAppPublisherValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/soft/updateAppPublisher',
                    JSON.stringify(vc.component.editAppPublisherInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppPublisherModel').modal('hide');
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
            refreshEditAppPublisherInfo: function() {
                vc.component.editAppPublisherInfo = {
                    publisherId: '',
                    username: '',
                    email: '',
                    phone: '',
                    extPublisherId:''
                }
            }
        }
    });

})(window.vc, window.vc.component);