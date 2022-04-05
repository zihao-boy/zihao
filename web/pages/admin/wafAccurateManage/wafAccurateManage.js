/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            wafCCManageInfo: {
                wafCCs: [],
                total: 0,
                records: 1,
                moreCondition: false,
                id: '',
                conditions: {
                    id: '',
                    blockSec: '',
                    path: '',

                }
            }
        },
        _initMethod: function() {
            vc.component._listWafCCs(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('wafCCManage', 'listWafCC', function(_param) {
                vc.component._listWafCCs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listWafCCs(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listWafCCs: function(_page, _rows) {

                vc.component.wafCCManageInfo.conditions.page = _page;
                vc.component.wafCCManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.wafCCManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafCC',
                    param,
                    function(json, res) {
                        var _wafCCManageInfo = JSON.parse(json);
                        vc.component.wafCCManageInfo.total = _wafCCManageInfo.total;
                        vc.component.wafCCManageInfo.records = _wafCCManageInfo.records;
                        vc.component.wafCCManageInfo.wafCCs = _wafCCManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.wafCCManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddWafCCModal: function() {
                vc.emit('addWafCC', 'openAddWafCCModal', {});
            },
            _openEditWafCCModel: function(_wafCC) {
                vc.emit('editWafCC', 'openEditWafCCModal', _wafCC);
            },
            _openDeleteWafCCModel: function(_wafCC) {
                vc.emit('deleteWafCC', 'openDeleteWafCCModal', _wafCC);
            },
            _queryWafCCMethod: function() {
                vc.component._listWafCCs(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.wafCCManageInfo.moreCondition) {
                    vc.component.wafCCManageInfo.moreCondition = false;
                } else {
                    vc.component.wafCCManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);