(function(vc, vm) {

    vc.extends({
        data: {
            deleteDbInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteDb', 'openDeleteDbModal', function(_params) {

                vc.component.deleteDbInfo = _params;
                $('#deleteDbModel').modal('show');

            });
        },
        methods: {
            deleteDb: function() {
                vc.http.apiPost(
                    '/resources/deleteDb',
                    JSON.stringify(vc.component.deleteDbInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteDbModel').modal('hide');
                            vc.emit('dbManage', 'listDb', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteDbModel: function() {
                $('#deleteDbModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);