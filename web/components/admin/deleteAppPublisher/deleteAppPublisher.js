(function(vc, vm) {

    vc.extends({
        data: {
            deleteAppPublisherInfo: {

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('deleteAppPublisher', 'openDeleteAppPublisherModal', function(_params) {

                vc.component.deleteAppPublisherInfo = _params;
                $('#deleteAppPublisherModel').modal('show');

            });
        },
        methods: {
            deleteAppPublisher: function() {
                vc.http.apiPost(
                    '/soft/deleteAppPublisher',
                    JSON.stringify(vc.component.deleteAppPublisherInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#deleteAppPublisherModel').modal('hide');
                            vc.emit('appPublisherManage', 'listAppPublisher', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.message(json);

                    });
            },
            closeDeleteAppPublisherModel: function() {
                $('#deleteAppPublisherModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);