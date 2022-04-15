/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            dnsMapManageInfo: {
                dnsMaps: [],
                total: 0,
                records: 1,
                moreCondition: false,
                dnsMapId: '',
                conditions: {
                    host: '',
                    type: '',
                    value: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listDnsMaps(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('dnsMapManage', 'listDnsMap', function (_param) {
                vc.component._listDnsMaps(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listDnsMaps(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listDnsMaps: function (_page, _rows) {

                vc.component.dnsMapManageInfo.conditions.page = _page;
                vc.component.dnsMapManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.dnsMapManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/dns/getDnsMap',
                    param,
                    function (json, res) {
                        var _dnsMapManageInfo = JSON.parse(json);
                        vc.component.dnsMapManageInfo.total = _dnsMapManageInfo.total;
                        vc.component.dnsMapManageInfo.records = _dnsMapManageInfo.records;
                        vc.component.dnsMapManageInfo.dnsMaps = _dnsMapManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.dnsMapManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddDnsMapModal: function () {
                vc.emit('addDnsMap', 'openAddDnsMapModal', {});
            },
            _openEditDnsMapModel: function (_dnsMap) {
                vc.emit('editDnsMap', 'openEditDnsMapModal', _dnsMap);
            },
            _openDeleteDnsMapModel: function (_dnsMap) {
                vc.emit('deleteDnsMap', 'openDeleteDnsMapModal', _dnsMap);
            },
            _queryDnsMapMethod: function () {
                vc.component._listDnsMaps(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.dnsMapManageInfo.moreCondition) {
                    vc.component.dnsMapManageInfo.moreCondition = false;
                } else {
                    vc.component.dnsMapManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
