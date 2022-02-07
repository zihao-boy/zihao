(function(vc, vm) {

    vc.extends({
        data: {
            uninstallImagesInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('uninstallImages', 'openUninstallImagesModal', function(_params) {

                vc.component.uninstallImagesInfo = _params;
                $('#uninstallImagesModel').modal('show');

            });
        },
        methods: {
            _uninstallImages: function() {
                vc.http.apiPost(
                    '/soft/uninstallImages',
                    JSON.stringify(vc.component.uninstallImagesInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#uninstallImagesModel').modal('hide');
                            vc.emit('imagesPoolManage', 'listBusinessImages', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(json);

                    });
            },
            closeUninstallImagesModel: function() {
                $('#uninstallImagesModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);