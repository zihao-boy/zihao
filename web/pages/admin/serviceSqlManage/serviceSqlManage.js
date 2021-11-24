/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            serviceSqlManageInfo: {
                serviceSqls: [],
                total: 0,
                records: 1,
                moreCondition: false,
                sqlId: '',
                conditions: {
                    sqlCode: '',
                    sqlId: '',
                    remark: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listServiceSqls(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('serviceSqlManage', 'listServiceSql', function (_param) {
                vc.component._listServiceSqls(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listServiceSqls(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listServiceSqls: function (_page, _rows) {

                vc.component.serviceSqlManageInfo.conditions.page = _page;
                vc.component.serviceSqlManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.serviceSqlManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/system/getServiceSqls',
                    param,
                    function (json, res) {
                        var _serviceSqlManageInfo = JSON.parse(json);
                        vc.component.serviceSqlManageInfo.total = _serviceSqlManageInfo.total;
                        vc.component.serviceSqlManageInfo.records = _serviceSqlManageInfo.records;
                        vc.component.serviceSqlManageInfo.serviceSqls = _serviceSqlManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.serviceSqlManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddServiceSqlModal: function () {
                vc.emit('addServiceSql', 'openAddServiceSqlModal', {});
            },
            _openEditServiceSqlModel: function (_serviceSql) {
                vc.emit('editServiceSql', 'openEditServiceSqlModal', _serviceSql);
            },
            _openDeleteServiceSqlModel: function (_serviceSql) {
                vc.emit('deleteServiceSql', 'openDeleteServiceSqlModal', _serviceSql);
            },
            _queryServiceSqlMethod: function () {
                vc.component._listServiceSqls(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.serviceSqlManageInfo.moreCondition) {
                    vc.component.serviceSqlManageInfo.moreCondition = false;
                } else {
                    vc.component.serviceSqlManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
