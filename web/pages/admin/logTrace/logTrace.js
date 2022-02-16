/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            logTraceInfo: {
                traces: [],
                total: 0,
                records: 1,
                moreCondition: false,
                mhgId: '',
                conditions: {
                    name: '',
                    traceId: '',
                    id: '',

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

                vc.component.logTraceInfo.conditions.page = _page;
                vc.component.logTraceInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.logTraceInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/monitor/getMonitorEvents',
                    param,
                    function (json, res) {
                        var _logTraceInfo = JSON.parse(json);
                        vc.component.logTraceInfo.total = _logTraceInfo.total;
                        vc.component.logTraceInfo.records = _logTraceInfo.records;
                        vc.component.logTraceInfo.traces = _logTraceInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.logTraceInfo.records,
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
                if (vc.component.logTraceInfo.moreCondition) {
                    vc.component.logTraceInfo.moreCondition = false;
                } else {
                    vc.component.logTraceInfo.moreCondition = true;
                }
            },
            
        }
    });
})(window.vc);
