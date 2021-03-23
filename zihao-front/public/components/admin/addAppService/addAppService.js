(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addAppServiceInfo: {
                asId: '',
                asName: '',
                asType: '',
                asDesc: '',
                asCount:'1'

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addAppService', 'openAddAppServiceModal', function () {
                $('#addAppServiceModel').modal('show');
            });
        },
        methods: {
            addAppServiceValidate() {
                return vc.validate.validate({
                    addAppServiceInfo: vc.component.addAppServiceInfo
                }, {
                    'addAppServiceInfo.asName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "应用名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "应用名称太长"
                        },
                    ],
                    'addAppServiceInfo.asType': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "服务类型不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "服务类型格式错误"
                        },
                    ],
                    'addAppServiceInfo.asCount': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "副本数不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "副本数不是有效数字"
                        },
                    ],
                    'addAppServiceInfo.asDesc': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "服务描述不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "512",
                            errInfo: "描述太长"
                        },
                    ],




                });
            },
            saveAppServiceInfo: function () {
                if (!vc.component.addAppServiceValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addAppServiceInfo);
                    $('#addAppServiceModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/appService/saveAppService',
                    JSON.stringify(vc.component.addAppServiceInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppServiceModel').modal('hide');
                            vc.component.clearAddAppServiceInfo();
                            vc.emit('appServiceManage', 'listAppService', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddAppServiceInfo: function () {
                vc.component.addAppServiceInfo = {
                    asName: '',
                    asType: '',
                    asDesc: '',
                    asCount:'1'

                };
            }
        }
    });

})(window.vc);
