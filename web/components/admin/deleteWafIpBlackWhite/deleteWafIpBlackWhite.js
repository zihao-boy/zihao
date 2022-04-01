(function(vc, vm) {

    vc.extends({
        data: {
            deleteWafIpBlackWhiteInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteWafIpBlackWhite', 'openDeleteWafIpBlackWhiteModal', function(_params) {

                vc.component.deleteWafIpBlackWhiteInfo = _params;
                $('#deleteWafIpBlackWhiteModel').modal('show');

            });
        },
        methods: {
            deleteWafIpBlackWhite: function() {
                vc.http.apiPost(
                    '/firewall/deleteWafIpBlackWhite',
                    JSON.stringify(vc.component.deleteWafIpBlackWhiteInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteWafIpBlackWhiteModel').modal('hide');
                            vc.emit('wafIpBlackWhiteManage', 'listWafIpBlackWhite', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteWafIpBlackWhiteModel: function() {
                $('#deleteWafIpBlackWhiteModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);