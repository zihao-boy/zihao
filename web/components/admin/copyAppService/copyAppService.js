(function (vc, vm) {

    vc.extends({
        data: {
            copyAppServiceInfo: {
                asId: '',
                asName: '',
                asDesc: '',
                imagesId: '',
                images: []

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('copyAppService', 'openCopyAppServiceModal', function (_params) {
                vc.component.refreshCopyAppServiceInfo();
                $that._listEditBusinessImagess();
                $('#copyAppServiceModel').modal('show');
                vc.copyObject(_params, vc.component.copyAppServiceInfo);
            });
        },
        methods: {
            copyAppServiceValidate: function () {
                return vc.validate.validate({
                    copyAppServiceInfo: vc.component.copyAppServiceInfo
                }, {
                    'copyAppServiceInfo.asName': [
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
                    'copyAppServiceInfo.asDesc': [
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
                    'copyAppServiceInfo.asId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "ID不能为空"
                        }]

                });
            },
            copyAppService: function () {
                if (!vc.component.copyAppServiceValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/appService/copyAppService',
                    JSON.stringify(vc.component.copyAppServiceInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#copyAppServiceModel').modal('hide');
                            vc.emit('appServiceControlInfo', 'notify', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshCopyAppServiceInfo: function () {
                vc.component.copyAppServiceInfo = {
                    asId: '',
                    asName: '',
                    asDesc: '',
                    imagesId: '',
                    images: []
                }
            },
            _listEditBusinessImagess: function (_page, _rows) {
                var param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImages',
                    param,
                    function (json, res) {
                        var _businessImagesManageInfo = JSON.parse(json);
                        vc.component.copyAppServiceInfo.images = _businessImagesManageInfo.data;
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc, window.vc.component);
