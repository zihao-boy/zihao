/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 20;
    vc.extends({
        data: {
            wafAccessLogInfo: {
                logs: [],
                total: 0,
                records: 1,
                moreCondition: false,
                mhgId: '',
                conditions: {
                    xRealIp: '',
                    scheme: '',
                    httpHost: '',

                }
            }
        },
        _initMethod: function() {
            vc.component._listWafAccessLogs(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('monitorHostGroupManage', 'listMonitorHostGroup', function(_param) {
                vc.component._listWafAccessLogs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listWafAccessLogs(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listWafAccessLogs: function(_page, _rows) {

                vc.component.wafAccessLogInfo.conditions.page = _page;
                vc.component.wafAccessLogInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.wafAccessLogInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafAccessLog',
                    param,
                    function(json, res) {
                        var _wafAccessLogInfo = JSON.parse(json);
                        vc.component.wafAccessLogInfo.total = _wafAccessLogInfo.total;
                        vc.component.wafAccessLogInfo.records = _wafAccessLogInfo.records;
                        vc.component.wafAccessLogInfo.traces = _wafAccessLogInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.wafAccessLogInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _queryMonitorHostGroupMethod: function() {
                vc.component._listWafAccessLogs(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function() {
                if (vc.component.wafAccessLogInfo.moreCondition) {
                    vc.component.wafAccessLogInfo.moreCondition = false;
                } else {
                    vc.component.wafAccessLogInfo.moreCondition = true;
                }
            },
            _toTraceDetail: function(_trace) {
                vc.jumpToPage('/index.html#/pages/admin/wafAccessLogDetail?traceId=' + _trace.traceId)
            }

        }
    });
})(window.vc);