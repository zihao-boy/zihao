(function(vc, vm) {

    vc.extends({
        data: {
            deleteWafAreaInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteWafArea', 'openDeleteWafAreaModal', function(_params) {

                vc.component.deleteWafAreaInfo = _params;
                $('#deleteWafAreaModel').modal('show');

            });
        },
        methods: {
            deleteWafArea: function() {
                vc.http.apiPost(
                    '/firewall/deleteWafArea',
                    JSON.stringify(vc.component.deleteWafAreaInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteWafAreaModel').modal('hide');
                            vc.emit('wafAreaManage', 'listWafArea', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteWafAreaModel: function() {
                $('#deleteWafAreaModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);