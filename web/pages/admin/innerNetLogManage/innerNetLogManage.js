/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            innerNetLogManageInfo: {
                innerNetLogs: [],
                total: 0,
                records: 1,
                moreCondition: false,
                pId: '',
                conditions: {
                    logId: '',
                    ip: '',
                    state: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listInnerNetLogs(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('innerNetLogManage', 'listInnerNetLog', function (_param) {
                vc.component._listInnerNetLogs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listInnerNetLogs(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listInnerNetLogs: function (_page, _rows) {

                vc.component.innerNetLogManageInfo.conditions.page = _page;
                vc.component.innerNetLogManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.innerNetLogManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/innerNet/getInnerNetLog',
                    param,
                    function (json, res) {
                        var _innerNetLogManageInfo = JSON.parse(json);
                        vc.component.innerNetLogManageInfo.total = _innerNetLogManageInfo.total;
                        vc.component.innerNetLogManageInfo.records = _innerNetLogManageInfo.records;
                        vc.component.innerNetLogManageInfo.innerNetLogs = _innerNetLogManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.innerNetLogManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _queryInnerNetLogMethod: function () {
                vc.component._listInnerNetLogs(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.innerNetLogManageInfo.moreCondition) {
                    vc.component.innerNetLogManageInfo.moreCondition = false;
                } else {
                    vc.component.innerNetLogManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
