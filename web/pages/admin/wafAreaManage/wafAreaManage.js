/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            wafAreaManageInfo: {
                wafAreas: [],
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
            vc.component._listWafAreas(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('wafAreaManage', 'listWafArea', function(_param) {
                vc.component._listWafAreas(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listWafAreas(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listWafAreas: function(_page, _rows) {

                vc.component.wafAreaManageInfo.conditions.page = _page;
                vc.component.wafAreaManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.wafAreaManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafArea',
                    param,
                    function(json, res) {
                        var _wafAreaManageInfo = JSON.parse(json);
                        vc.component.wafAreaManageInfo.total = _wafAreaManageInfo.total;
                        vc.component.wafAreaManageInfo.records = _wafAreaManageInfo.records;
                        vc.component.wafAreaManageInfo.wafAreas = _wafAreaManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.wafAreaManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddWafAreaModal: function() {
                vc.emit('addWafArea', 'openAddWafAreaModal', {});
            },
            _openEditWafAreaModel: function(_wafArea) {
                vc.emit('editWafArea', 'openEditWafAreaModal', _wafArea);
            },
            _openDeleteWafAreaModel: function(_wafArea) {
                vc.emit('deleteWafArea', 'openDeleteWafAreaModal', _wafArea);
            },
            _queryWafAreaMethod: function() {
                vc.component._listWafAreas(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.wafAreaManageInfo.moreCondition) {
                    vc.component.wafAreaManageInfo.moreCondition = false;
                } else {
                    vc.component.wafAreaManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);