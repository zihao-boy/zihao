(function(vc, vm) {

    vc.extends({
        data: {
            deleteOssInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteOss', 'openDeleteOssModal', function(_params) {

                vc.component.deleteOssInfo = _params;
                $('#deleteOssModel').modal('show');

            });
        },
        methods: {
            deleteOss: function() {
                vc.http.apiPost(
                    '/resources/deleteOss',
                    JSON.stringify(vc.component.deleteOssInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteOssModel').modal('hide');
                            vc.emit('ossManage', 'listOss', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteOssModel: function() {
                $('#deleteOssModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);