/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            monitorHostManageInfo: {
                monitorHosts: [],
                total: 0,
                records: 1,
                moreCondition: false,
                mhId: '',
                conditions: {
                    name: '',
                    mhgId: vc.getParam('mhgId')
                }
            }
        },
        _initMethod: function () {
            vc.component._listMonitorHosts(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('monitorHostManage', 'listMonitorHost', function (_param) {
                vc.component._listMonitorHosts(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listMonitorHosts(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listMonitorHosts: function (_page, _rows) {

                vc.component.monitorHostManageInfo.conditions.page = _page;
                vc.component.monitorHostManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.monitorHostManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/monitor/getMonitorHosts',
                    param,
                    function (json, res) {
                        var _monitorHostManageInfo = JSON.parse(json);
                        vc.component.monitorHostManageInfo.total = _monitorHostManageInfo.total;
                        vc.component.monitorHostManageInfo.records = _monitorHostManageInfo.records;
                        vc.component.monitorHostManageInfo.monitorHosts = _monitorHostManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.monitorHostManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddMonitorHostModal: function () {
                vc.emit('addMonitorHost', 'openAddMonitorHostModal', {
                    mhgId: vc.getParam('mhgId')
                });
            },
            _openEditMonitorHostModel: function (_monitorHost) {
                vc.emit('editMonitorHost', 'openEditMonitorHostModal', _monitorHost);
            },
            _openDeleteMonitorHostModel: function (_monitorHost) {
                vc.emit('deleteMonitorHost', 'openDeleteMonitorHostModal', _monitorHost);
            },
            _queryMonitorHostMethod: function () {
                vc.component._listMonitorHosts(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.monitorHostManageInfo.moreCondition) {
                    vc.component.monitorHostManageInfo.moreCondition = false;
                } else {
                    vc.component.monitorHostManageInfo.moreCondition = true;
                }
            },
            _goBack: function () {
                vc.goBack();
            },
            conver: function (limit) {
                if(limit == null || limit == undefined || limit == ''){
                    return "0GB";
                }
                var size = "";
                if (limit < 0.1 * 1024) { //如果小于0.1KB转化成B  
                    size = limit.toFixed(2) + "MB";
                } else if (limit < 0.1 * 1024 * 1024) {//如果小于0.1MB转化成KB  
                    size = (limit / 1024).toFixed(2) + "GB";
                }

                var sizestr = size + "";
                var len = sizestr.indexOf("\.");
                var dec = sizestr.substr(len + 1, 2);
                if (dec == "00") {//当小数点后为00时 去掉小数部分  
                    return sizestr.substring(0, len) + sizestr.substr(len + 3, 2);
                }
                return sizestr;
            }


        }
    });
})(window.vc);
