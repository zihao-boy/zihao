(function(vc) {

    vc.extends({
        data: {
            editAppServicePortInfo: {
                asId: '',
                srcPort: '',
                targetPort: '',
                portId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editAppServicePort', 'openEditAppServicePortModal', function(_param) {
                vc.copyObject(_param, $that.editAppServicePortInfo);
                $('#editAppServicePortModel').modal('show');
            });
        },
        methods: {
            editAppServicePortValidate() {
                return vc.validate.validate({
                    editAppServicePortInfo: vc.component.editAppServicePortInfo
                }, {
                    'editAppServicePortInfo.srcPort': [{
                            limit: "required",
                            param: "",
                            errInfo: "宿主机端口不能为空"
                        }
                    ],
                    'editAppServicePortInfo.targetPort': [{
                            limit: "required",
                            param: "",
                            errInfo: "容器端口不能为空"
                        }
                    ]
                });
            },
            updateAppServicePortInfo: function() {
                if (!vc.component.editAppServicePortValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appService/updateAppServicePort',
                    JSON.stringify(vc.component.editAppServicePortInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppServicePortModel').modal('hide');
                            vc.emit('appServiceControlPort', 'switch', $that.editAppServicePortInfo);
                            vc.component.cleareditAppServicePortInfo();
                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            cleareditAppServicePortInfo: function() {
                vc.component.editAppServicePortInfo = {
                    asId: '',
                    srcPort: '',
                    targetPort: '',
                    portId: ''
                };
            }
        }
    });

})(window.vc);