(function(vc, vm) {

    vc.extends({
        data: {
            editFasterDeplyInfo: {
                shellPackageId: '',
                shellContext: '',
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editFasterDeply', 'openEditFasterDeplyModal', function(_params) {
                vc.component.refreshEditFasterDeplyInfo();
                $that.loadShellContext();
                $('#editFasterDeplyModel').modal('show');
                vc.copyObject(_params, vc.component.editFasterDeplyInfo);
            });
        },
        methods: {
            editFasterDeply: function() {
                vc.http.apiPost(
                    '/soft/editBusinessPackageContext',
                    JSON.stringify(vc.component.editFasterDeplyInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editFasterDeplyModel').modal('hide');
                            vc.emit('hostManage', 'listFasterDeply', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditFasterDeplyInfo: function() {
                vc.component.editFasterDeplyInfo = {
                    shellPackageId: '',
                    shellContext: '',
                }
            },
            loadShellContext: function() {
                var param = {
                    params: {
                        page: 1,
                        row: 1,
                        shellPackageId: $that.editFasterDeplyInfo.shellPackageId
                    }
                };

                //发送get请求
                vc.http.apiGet('/soft/listBusinessPackageContext',
                    param,
                    function(json, res) {
                        var _hostGroupManageInfo = JSON.parse(json);
                        vc.component.editFasterDeplyInfo.shellContext = _hostGroupManageInfo.data;

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }
        }
    });

})(window.vc, window.vc.component);