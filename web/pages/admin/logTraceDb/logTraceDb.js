/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            logTraceDbInfo: {
                traces: [],
                total: 0,
                records: 1,
                moreCondition: false,
                mhgId: '',
                conditions: {
                    serviceName: '',
                    traceId: '',
                    spanId: '',
                }
            }
        },
        _initMethod: function() {
            $that.logTraceDbInfo.conditions.spanId = vc.getParam('spanId')
            vc.component._listLogTraceDbs(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('monitorHostGroupManage', 'listMonitorHostGroup', function(_param) {
                vc.component._listLogTraceDbs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listLogTraceDbs(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listLogTraceDbs: function(_page, _rows) {

                vc.component.logTraceDbInfo.conditions.page = _page;
                vc.component.logTraceDbInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.logTraceDbInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/monitor/getLogTraceDb',
                    param,
                    function(json, res) {
                        var _logTraceDbInfo = JSON.parse(json);
                        vc.component.logTraceDbInfo.total = _logTraceDbInfo.total;
                        vc.component.logTraceDbInfo.records = _logTraceDbInfo.records;
                        vc.component.logTraceDbInfo.traces = _logTraceDbInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.logTraceDbInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _queryMonitorHostGroupMethod: function() {
                vc.component._listLogTraceDbs(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function() {
                if (vc.component.logTraceDbInfo.moreCondition) {
                    vc.component.logTraceDbInfo.moreCondition = false;
                } else {
                    vc.component.logTraceDbInfo.moreCondition = true;
                }
            },
            _viewSql: function(_trace) {
                vc.emit('viewLogTraceDb', 'openViewLogTraceDbModal', _trace);
            },
            _goBack: function() {
                vc.goBack();
            }

        }
    });
})(window.vc);