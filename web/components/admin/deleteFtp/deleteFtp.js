(function(vc, vm) {

    vc.extends({
        data: {
            deleteFtpInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteFtp', 'openDeleteFtpModal', function(_params) {

                vc.component.deleteFtpInfo = _params;
                $('#deleteFtpModel').modal('show');

            });
        },
        methods: {
            deleteFtp: function() {
                vc.http.apiPost(
                    '/resources/deleteFtp',
                    JSON.stringify(vc.component.deleteFtpInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteFtpModel').modal('hide');
                            vc.emit('ftpManage', 'listFtp', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteFtpModel: function() {
                $('#deleteFtpModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);