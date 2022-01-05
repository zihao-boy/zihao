(function(vc, vm) {

    vc.extends({
        data: {
            editFasterDeplyInfo: {
                hostId: '',
                groupId: '',


            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editFasterDeply', 'openEditFasterDeplyModal', function(_params) {
                vc.component.refreshEditFasterDeplyInfo();
                // $that.loadEditFasterDeplyFasterDeplyGroup();
                $('#editFasterDeplyModel').modal('show');
                vc.copyObject(_params, vc.component.editFasterDeplyInfo);
            });
        },
        methods: {
            editFasterDeply: function() {
                vc.http.apiPost(
                    '/host/updateFasterDeply',
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
                    hostId: '',
                    groupId: '',
                    name: '',
                    ip: '',
                    username: '',
                    passwd: '',
                    hostGroups: []
                }
            },
            loadEditFasterDeplyFasterDeplyGroup: function() {
                var param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getFasterDeplyGroup',
                    param,
                    function(json, res) {
                        var _hostGroupManageInfo = JSON.parse(json);
                        vc.component.editFasterDeplyInfo.hostGroups = _hostGroupManageInfo.data;

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }
        }
    });

})(window.vc, window.vc.component);