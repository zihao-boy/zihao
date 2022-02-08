(function(vc, vm) {

    vc.extends({
        data: {
            offlineApplyAppInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('offlineApplyApp', 'openOfflineApplyAppModal', function(_params) {

                vc.component.offlineApplyAppInfo = _params;
                $('#offlineApplyAppModel').modal('show');

            });
        },
        methods: {
            _offlineApplyApp: function() {
                vc.http.apiPost(
                    '/soft/offlineApplyApp',
                    JSON.stringify(vc.component.offlineApplyAppInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#offlineApplyAppModel').modal('hide');
                            vc.emit('myAppPoolManage', 'listBusinessImages', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(json);

                    });
            },
            closeOfflineApplyAppModel: function() {
                $('#offlineApplyAppModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);