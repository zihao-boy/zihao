/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            wafIpBlackWhiteManageInfo: {
                wafIpBlackWhites: [],
                total: 0,
                records: 1,
                moreCondition: false,
                id: '',
                conditions: {
                    id: '',
                    typeCd: '',
                    ip: '',

                }
            }
        },
        _initMethod: function() {
            vc.component._listWafIpBlackWhites(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('wafIpBlackWhiteManage', 'listWafIpBlackWhite', function(_param) {
                vc.component._listWafIpBlackWhites(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listWafIpBlackWhites(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listWafIpBlackWhites: function(_page, _rows) {

                vc.component.wafIpBlackWhiteManageInfo.conditions.page = _page;
                vc.component.wafIpBlackWhiteManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.wafIpBlackWhiteManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafIpBlackWhite',
                    param,
                    function(json, res) {
                        var _wafIpBlackWhiteManageInfo = JSON.parse(json);
                        vc.component.wafIpBlackWhiteManageInfo.total = _wafIpBlackWhiteManageInfo.total;
                        vc.component.wafIpBlackWhiteManageInfo.records = _wafIpBlackWhiteManageInfo.records;
                        vc.component.wafIpBlackWhiteManageInfo.wafIpBlackWhites = _wafIpBlackWhiteManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.wafIpBlackWhiteManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddWafIpBlackWhiteModal: function() {
                vc.emit('addWafIpBlackWhite', 'openAddWafIpBlackWhiteModal', {});
            },
            _openEditWafIpBlackWhiteModel: function(_wafIpBlackWhite) {
                vc.emit('editWafIpBlackWhite', 'openEditWafIpBlackWhiteModal', _wafIpBlackWhite);
            },
            _openDeleteWafIpBlackWhiteModel: function(_wafIpBlackWhite) {
                vc.emit('deleteWafIpBlackWhite', 'openDeleteWafIpBlackWhiteModal', _wafIpBlackWhite);
            },
            _queryWafIpBlackWhiteMethod: function() {
                vc.component._listWafIpBlackWhites(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.wafIpBlackWhiteManageInfo.moreCondition) {
                    vc.component.wafIpBlackWhiteManageInfo.moreCondition = false;
                } else {
                    vc.component.wafIpBlackWhiteManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);