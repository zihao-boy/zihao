/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostEventsInfo: {
                hostEvents: [],
                total: 0,
                records: 1,
                hostId: '',
            }
        },
        _initMethod: function () {
        },
        _initEvent: function () {

            vc.on('hostEvents', 'switch', function (_param) {
                if (_param.hostId == '') {
                    return;
                }
                vc.copyObject(_param, $that.hostEventsInfo)
                vc.component._listhostEvents(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listhostEvents(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listhostEvents: function (_page, _rows) {

                var param = {
                    params: {
                        page:_page,
                        row:_rows,
                        eventObjId:$that.hostEventsInfo.hostId
                    }
                };

                //发送get请求
                vc.http.apiGet('/monitor/getMonitorEvents',
                    param,
                    function (json, res) {
                        var _hostEventsInfo = JSON.parse(json);
                        vc.component.hostEventsInfo.total = _hostEventsInfo.total;
                        vc.component.hostEventsInfo.records = _hostEventsInfo.records;
                        vc.component.hostEventsInfo.hostEvents = _hostEventsInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.hostEventsInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }
            
        }
    });
})(window.vc);
