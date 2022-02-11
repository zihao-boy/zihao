(function(vc, vm) {

    vc.extends({
        data: {
            deleteTenantSettingInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteTenantSetting', 'openDeleteTenantSettingModal', function(_params) {

                vc.component.deleteTenantSettingInfo = _params;
                $('#deleteTenantSettingModel').modal('show');

            });
        },
        methods: {
            deleteTenantSetting: function() {
                vc.http.apiPost(
                    'tenantSetting.deleteTenantSetting',
                    JSON.stringify(vc.component.deleteTenantSettingInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteTenantSettingModel').modal('hide');
                            vc.emit('tenantSettingManage', 'listTenantSetting', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteTenantSettingModel: function() {
                $('#deleteTenantSettingModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);