(function(vc, vm) {

    vc.extends({
        data: {
            upgradeAppServiceInfo: {
                asId: '',
                asName: '',
                asDesc: '',
                imagesId: '',
                verId: '',
                images: [],
                vers: [],
                asCount: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('upgradeAppService', 'openUpgradeAppServiceModal', function(_params) {
                vc.component.refreshUpgradeAppServiceInfo();
                $that._listUpgradeBusinessImagess();
                $('#upgradeAppServiceModel').modal('show');
                vc.copyObject(_params, vc.component.upgradeAppServiceInfo);
                $that._upgradeChangeImages();
            });
        },
        methods: {
            upgradeAppServiceValidate: function() {
                return vc.validate.validate({
                    upgradeAppServiceInfo: vc.component.upgradeAppServiceInfo
                }, {
                    'upgradeAppServiceInfo.asName': [{
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
                    'upgradeAppServiceInfo.asId': [{
                        limit: "required",
                        param: "",
                        errInfo: "ID不能为空"
                    }]

                });
            },
            upgradeAppService: function() {
                if (!vc.component.upgradeAppServiceValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/appService/upgradeAppService',
                    JSON.stringify(vc.component.upgradeAppServiceInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#upgradeAppServiceModel').modal('hide');
                            vc.emit('appServiceManage', 'listAppService', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshUpgradeAppServiceInfo: function() {
                vc.component.upgradeAppServiceInfo = {
                    asId: '',
                    asName: '',
                    asDesc: '',
                    imagesId: '',
                    verId: '',
                    images: [],
                    vers: [],
                    asCount: '',
                }
            },
            _listUpgradeBusinessImagess: function(_page, _rows) {
                var param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImages',
                    param,
                    function(json, res) {
                        var _businessImagesManageInfo = JSON.parse(json);
                        vc.component.upgradeAppServiceInfo.images = _businessImagesManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _upgradeChangeImages: function() {
                var param = {
                    params: {
                        page: 1,
                        row: 50,
                        imagesId: $that.upgradeAppServiceInfo.imagesId
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/getBusinessImagesVer',
                    param,
                    function(json) {
                        var _verInfo = JSON.parse(json);
                        vc.component.upgradeAppServiceInfo.vers = _verInfo.data;

                    },
                    function() {
                        console.log('请求失败处理');
                    }
                );
            }
        }
    });

})(window.vc, window.vc.component);