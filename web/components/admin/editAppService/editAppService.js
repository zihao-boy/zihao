(function (vc, vm) {

    vc.extends({
        data: {
            editAppServiceInfo: {
                asId: '',
                asName: '',
                asType: '',
                asDesc: '',
                asCount:''

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editAppService', 'openEditAppServiceModal', function (_params) {
                vc.component.refreshEditAppServiceInfo();
                $('#editAppServiceModel').modal('show');
                vc.copyObject(_params, vc.component.editAppServiceInfo);
            });
        },
        methods: {
            editAppServiceValidate: function () {
                return vc.validate.validate({
                    editAppServiceInfo: vc.component.editAppServiceInfo
                }, {
                    'editAppServiceInfo.asName': [
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
                    'editAppServiceInfo.asType': [
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
                    'editAppServiceInfo.asCount': [
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
                    'editAppServiceInfo.asDesc': [
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
                    'editAppServiceInfo.asId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "ID不能为空"
                        }]

                });
            },
            editAppService: function () {
                if (!vc.component.editAppServiceValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/appService/updateAppService',
                    JSON.stringify(vc.component.editAppServiceInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppServiceModel').modal('hide');
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
            refreshEditAppServiceInfo: function () {
                vc.component.editAppServiceInfo = {
                    asId: '',
                    asName: '',
                    asType: '',
                    asDesc: '',
                    asCount:''

                }
            }
        }
    });

})(window.vc, window.vc.component);
