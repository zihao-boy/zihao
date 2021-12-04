(function(vc, vm) {

    vc.extends({
        data: {
            deleteBusinessDockerfileInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteBusinessDockerfile', 'openDeleteBusinessDockerfileModal', function(_params) {

                vc.component.deleteBusinessDockerfileInfo = _params;
                $('#deleteBusinessDockerfileModel').modal('show');
            });
        },
        methods: {
            deleteBusinessDockerfile: function() {
                vc.http.apiPost(
                    '/soft/deleteBusinessDockerfile',
                    JSON.stringify(vc.component.deleteBusinessDockerfileInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteBusinessDockerfileModel').modal('hide');
                            vc.emit('businessDockerfileManage', 'listBusinessDockerfile', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteBusinessDockerfileModel: function() {
                $('#deleteBusinessDockerfileModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);