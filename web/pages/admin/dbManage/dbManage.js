/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            dbManageInfo: {
                dbs: [],
                total: 0,
                records: 1,
                moreCondition: false,
                dbId: '',
                conditions: {
                    dbId: '',
                    name: '',
                    ip: '',

                }
            }
        },
        _initMethod: function() {
            vc.component._listDbs(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('dbManage', 'listDb', function(_param) {
                vc.component._listDbs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listDbs(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listDbs: function(_page, _rows) {

                vc.component.dbManageInfo.conditions.page = _page;
                vc.component.dbManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.dbManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/resources/getDb',
                    param,
                    function(json, res) {
                        var _dbManageInfo = JSON.parse(json);
                        vc.component.dbManageInfo.total = _dbManageInfo.total;
                        vc.component.dbManageInfo.records = _dbManageInfo.records;
                        vc.component.dbManageInfo.dbs = _dbManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.dbManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddDbModal: function() {
                vc.emit('addDb', 'openAddDbModal', {});
            },
            _openEditDbModel: function(_db) {
                vc.emit('editDb', 'openEditDbModal', _db);
            },
            _openDeleteDbModel: function(_db) {
                vc.emit('deleteDb', 'openDeleteDbModal', _db);
            },
            _queryDbMethod: function() {
                vc.component._listDbs(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.dbManageInfo.moreCondition) {
                    vc.component.dbManageInfo.moreCondition = false;
                } else {
                    vc.component.dbManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);