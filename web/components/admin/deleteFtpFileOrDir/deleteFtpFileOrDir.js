(function(vc, vm) {

    vc.extends({
        data: {
            deleteFileOrDirInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteFileOrDir', 'openDeleteFileOrDirModal', function(_params) {

                vc.component.deleteFileOrDirInfo = _params;
                $('#deleteFileOrDirModel').modal('show');

            });
        },
        methods: {
            deleteFileOrDir: function() {
                vc.http.apiPost(
                    '/resources/removeFtpFile',
                    JSON.stringify(vc.component.deleteFileOrDirInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        //let data = res.data;
                        if (_json.code == 200 || _json.code == 0) {
                            //关闭model
                            $('#deleteFileOrDirModel').modal('hide');
                            vc.component.closeDeleteFileOrDirModel();
                            vc.emit('ftpFileManager', 'listFiles', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(json);

                    });
            },
            closeDeleteFileOrDirModel: function() {
                $('#deleteFileOrDirModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);