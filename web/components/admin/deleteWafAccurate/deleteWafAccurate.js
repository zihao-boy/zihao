(function(vc, vm) {

    vc.extends({
        data: {
            deleteWafAccurateInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteWafAccurate', 'openDeleteWafAccurateModal', function(_params) {

                vc.component.deleteWafAccurateInfo = _params;
                $('#deleteWafAccurateModel').modal('show');

            });
        },
        methods: {
            deleteWafAccurate: function() {
                vc.http.apiPost(
                    '/firewall/deleteWafAccurate',
                    JSON.stringify(vc.component.deleteWafAccurateInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteWafAccurateModel').modal('hide');
                            vc.emit('wafAccurateManage', 'listWafAccurate', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteWafAccurateModel: function() {
                $('#deleteWafAccurateModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);