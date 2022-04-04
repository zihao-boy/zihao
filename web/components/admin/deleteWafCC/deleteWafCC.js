(function(vc, vm) {

    vc.extends({
        data: {
            deleteWafCCInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteWafCC', 'openDeleteWafCCModal', function(_params) {

                vc.component.deleteWafCCInfo = _params;
                $('#deleteWafCCModel').modal('show');

            });
        },
        methods: {
            deleteWafCC: function() {
                vc.http.apiPost(
                    '/firewall/deleteWafCC',
                    JSON.stringify(vc.component.deleteWafCCInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteWafCCModel').modal('hide');
                            vc.emit('wafCCManage', 'listWafCC', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteWafCCModel: function() {
                $('#deleteWafCCModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);