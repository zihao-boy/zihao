/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            monitorEventManageInfo: {
                monitorEvents: [],
                total: 0,
                records: 1,
                moreCondition: false,
                mhgId: '',
                conditions: {
                    name: '',
                    state: '',
                    noticeType: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listmonitorEvents(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('monitorHostGroupManage', 'listMonitorHostGroup', function (_param) {
                vc.component._listmonitorEvents(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listmonitorEvents(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listmonitorEvents: function (_page, _rows) {

                vc.component.monitorEventManageInfo.conditions.page = _page;
                vc.component.monitorEventManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.monitorEventManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/monitor/getMonitorEvents',
                    param,
                    function (json, res) {
                        var _monitorEventManageInfo = JSON.parse(json);
                        vc.component.monitorEventManageInfo.total = _monitorEventManageInfo.total;
                        vc.component.monitorEventManageInfo.records = _monitorEventManageInfo.records;
                        vc.component.monitorEventManageInfo.monitorEvents = _monitorEventManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.monitorEventManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _queryMonitorHostGroupMethod: function () {
                vc.component._listmonitorEvents(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function () {
                if (vc.component.monitorEventManageInfo.moreCondition) {
                    vc.component.monitorEventManageInfo.moreCondition = false;
                } else {
                    vc.component.monitorEventManageInfo.moreCondition = true;
                }
            },
            
        }
    });
})(window.vc);
