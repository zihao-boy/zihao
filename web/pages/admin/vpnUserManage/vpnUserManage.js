/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            vpnUserManageInfo: {
                vpnUsers: [],
                total: 0,
                records: 1,
                moreCondition: false,
                routeId: '',
                conditions: {
                    routeId: '',
                    hostname: '',
                    ip: '',
                    port: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listVpnUsers(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('vpnUserManage', 'listVpnUser', function (_param) {
                vc.component._listVpnUsers(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listVpnUsers(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listVpnUsers: function (_page, _rows) {

                vc.component.vpnUserManageInfo.conditions.page = _page;
                vc.component.vpnUserManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.vpnUserManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/vpn/getVpnUser',
                    param,
                    function (json, res) {
                        var _vpnUserManageInfo = JSON.parse(json);
                        vc.component.vpnUserManageInfo.total = _vpnUserManageInfo.total;
                        vc.component.vpnUserManageInfo.records = _vpnUserManageInfo.records;
                        vc.component.vpnUserManageInfo.vpnUsers = _vpnUserManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.vpnUserManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddVpnUserModal: function () {
                vc.emit('addVpnUser', 'openAddVpnUserModal', {});
            },
            _openEditVpnUserModel: function (_vpnUser) {
                vc.emit('editVpnUser', 'openEditVpnUserModal', _vpnUser);
            },
            _openDeleteVpnUserModel: function (_vpnUser) {
                vc.emit('deleteVpnUser', 'openDeleteVpnUserModal', _vpnUser);
            },
            _queryVpnUserMethod: function () {
                vc.component._listVpnUsers(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.vpnUserManageInfo.moreCondition) {
                    vc.component.vpnUserManageInfo.moreCondition = false;
                } else {
                    vc.component.vpnUserManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
