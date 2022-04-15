/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            dnsManageInfo: {
                dnss: [],
                total: 0,
                records: 1,
                moreCondition: false,
                dnsId: '',
                conditions: {
                    dnsId: '',
                    dnsName: '',
                    state: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listDnss(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('dnsManage', 'listDns', function (_param) {
                vc.component._listDnss(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listDnss(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listDnss: function (_page, _rows) {

                vc.component.dnsManageInfo.conditions.page = _page;
                vc.component.dnsManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.dnsManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/dns/getDns',
                    param,
                    function (json, res) {
                        var _dnsManageInfo = JSON.parse(json);
                        vc.component.dnsManageInfo.total = _dnsManageInfo.total;
                        vc.component.dnsManageInfo.records = _dnsManageInfo.records;
                        vc.component.dnsManageInfo.dnss = _dnsManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.dnsManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openEditDnsModel: function (_dns) {
                vc.emit('editDns', 'openEditDnsModal', _dns);
            },
            _openStartDnsModel: function (_dns) {
                vc.http.apiPost(
                    '/dns/startDns',
                    JSON.stringify(_dns),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('dnsManage', 'listDns', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            _openStopDnsModel: function (_dns) {
                vc.http.apiPost(
                    '/dns/stopDns',
                    JSON.stringify(_dns),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('dnsManage', 'listDns', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            _queryDnsMethod: function () {
                vc.component._listDnss(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.dnsManageInfo.moreCondition) {
                    vc.component.dnsManageInfo.moreCondition = false;
                } else {
                    vc.component.dnsManageInfo.moreCondition = true;
                }
            },
            _loadDnsConfig: function (_dns) {
                vc.http.apiPost(
                    '/dns/refreshDnsConfig',
                    JSON.stringify(_dns),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.emit('dnsManage', 'listDns', {});
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
