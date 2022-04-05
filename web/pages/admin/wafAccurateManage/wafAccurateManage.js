/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            wafAccurateManageInfo: {
                wafAccurates: [],
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
            vc.component._listWafAccurates(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('wafAccurateManage', 'listWafAccurate', function(_param) {
                vc.component._listWafAccurates(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listWafAccurates(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listWafAccurates: function(_page, _rows) {

                vc.component.wafAccurateManageInfo.conditions.page = _page;
                vc.component.wafAccurateManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.wafAccurateManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafAccurate',
                    param,
                    function(json, res) {
                        var _wafAccurateManageInfo = JSON.parse(json);
                        vc.component.wafAccurateManageInfo.total = _wafAccurateManageInfo.total;
                        vc.component.wafAccurateManageInfo.records = _wafAccurateManageInfo.records;
                        vc.component.wafAccurateManageInfo.wafAccurates = _wafAccurateManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.wafAccurateManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddWafAccurateModal: function() {
                vc.emit('addWafAccurate', 'openAddWafAccurateModal', {});
            },
            _openEditWafAccurateModel: function(_wafAccurate) {
                vc.emit('editWafAccurate', 'openEditWafAccurateModal', _wafAccurate);
            },
            _openDeleteWafAccurateModel: function(_wafAccurate) {
                vc.emit('deleteWafAccurate', 'openDeleteWafAccurateModal', _wafAccurate);
            },
            _queryWafAccurateMethod: function() {
                vc.component._listWafAccurates(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.wafAccurateManageInfo.moreCondition) {
                    vc.component.wafAccurateManageInfo.moreCondition = false;
                } else {
                    vc.component.wafAccurateManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);