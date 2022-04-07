/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            vpnManageInfo: {
                vpns: [],
                total: 0,
                records: 1,
                moreCondition: false,
                vpnId: '',
                conditions: {
                    vpnId: '',
                    vpnName: '',
                    state: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listVpns(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('vpnManage', 'listVpn', function (_param) {
                vc.component._listVpns(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listVpns(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listVpns: function (_page, _rows) {

                vc.component.vpnManageInfo.conditions.page = _page;
                vc.component.vpnManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.vpnManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/vpn/getVpn',
                    param,
                    function (json, res) {
                        var _vpnManageInfo = JSON.parse(json);
                        vc.component.vpnManageInfo.total = _vpnManageInfo.total;
                        vc.component.vpnManageInfo.records = _vpnManageInfo.records;
                        vc.component.vpnManageInfo.vpns = _vpnManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.vpnManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openEditVpnModel: function (_vpn) {
                vc.emit('editVpn', 'openEditVpnModal', _vpn);
            },
            _openStartVpnModel: function (_vpn) {
                vc.http.apiPost(
                    '/vpn/startVpn',
                    JSON.stringify(_vpn),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('vpnManage', 'listVpn', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            _openStopVpnModel: function (_vpn) {
                vc.http.apiPost(
                    '/vpn/stopVpn',
                    JSON.stringify(_vpn),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('vpnManage', 'listVpn', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            _queryVpnMethod: function () {
                vc.component._listVpns(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.vpnManageInfo.moreCondition) {
                    vc.component.vpnManageInfo.moreCondition = false;
                } else {
                    vc.component.vpnManageInfo.moreCondition = true;
                }
            },
            _loadVpnConfig: function (_vpn) {
                vc.http.apiPost(
                    '/vpn/refreshVpnConfig',
                    JSON.stringify(_vpn),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('vpnManage', 'listVpn', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
           


        }
    });
})(window.vc);
