(function(vc, vm) {
    vc.extends({
        data: {
            deleteFasterDeployInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteFasterDeploy', 'openDeleteFasterDeployModal', function(_params) {
                vc.component.deleteFasterDeployInfo = _params;
                $('#deleteFasterDeployModel').modal('show');
            });
        },
        methods: {
            deleteFasterDeploy: function() {
                vc.http.apiPost(
                    '/appService/deleteFasterDeploy',
                    JSON.stringify(vc.component.deleteFasterDeployInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteFasterDeployModel').modal('hide');
                            vc.emit('fasterDeployLog', 'listFasterDeployLog', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(json);

                    });
            },
            closeDeleteFasterDeployModel: function() {
                $('#deleteFasterDeployModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);