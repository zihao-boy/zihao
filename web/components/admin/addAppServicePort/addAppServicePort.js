(function(vc) {

    vc.extends({
        data: {
            addAppServicePortInfo: {
                asId: '',
                srcPort: '',
                targetPort: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addAppServicePort', 'openAddAppServicePortModal', function(_param) {
                $that.addAppServicePortInfo.asId = _param.asId;
                $('#addAppServicePortModel').modal('show');
            });
        },
        methods: {
            addAppServicePortValidate() {
                return vc.validate.validate({
                    addAppServicePortInfo: vc.component.addAppServicePortInfo
                }, {
                    'addAppServicePortInfo.srcPort': [{
                            limit: "required",
                            param: "",
                            errInfo: "宿主机端口不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "宿主机端口为正数"
                        },
                    ],
                    'addAppServicePortInfo.targetPort': [{
                            limit: "required",
                            param: "",
                            errInfo: "容器端口不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "容器端口为正数"
                        },
                    ]
                });
            },
            saveAppServicePortInfo: function() {
                if (!vc.component.addAppServicePortValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appService/saveAppServicePort',
                    JSON.stringify(vc.component.addAppServicePortInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppServicePortModel').modal('hide');
                            vc.emit('appServiceControlPort', 'switch', $that.addAppServicePortInfo);
                            vc.component.clearaddAppServicePortInfo();
                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearaddAppServicePortInfo: function() {
                vc.component.addAppServicePortInfo = {
                    asId: '',
                    srcPort: '',
                    targetPort: ''
                };
            }
        }
    });

})(window.vc);