(function(vc, vm) {

    vc.extends({
        data: {
            deleteBackupInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteBackup', 'openDeleteBackupModal', function(_params) {

                vc.component.deleteBackupInfo = _params;
                $('#deleteBackupModel').modal('show');

            });
        },
        methods: {
            deleteBackup: function() {
                vc.http.apiPost(
                    '/resources/deleteBackUp',
                    JSON.stringify(vc.component.deleteBackupInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteBackupModel').modal('hide');
                            vc.emit('backupManage', 'listBackup', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteBackupModel: function() {
                $('#deleteBackupModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);