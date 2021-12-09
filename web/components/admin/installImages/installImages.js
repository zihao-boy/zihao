(function(vc, vm) {

    vc.extends({
        data: {
            installImagesInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('installImages', 'openInstallImagesModal', function(_params) {

                vc.component.installImagesInfo = _params;
                $('#installImagesModel').modal('show');

            });
        },
        methods: {
            _installImages: function() {
                vc.http.apiPost(
                    '/soft/installImages',
                    JSON.stringify(vc.component.installImagesInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#installImagesModel').modal('hide');
                            vc.jumpToPage('/index.html#/pages/admin/businessImagesManage')
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(json);

                    });
            },
            closeInstallImagesModel: function() {
                $('#installImagesModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);