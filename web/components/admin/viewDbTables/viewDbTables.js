(function (vc, vm) {

    vc.extends({
        data: {
            viewDbTablesInfo: {
                curDbId: '',
                curDbName: '',
                tables: []
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('viewDbTables', 'openViewDbDataModal', function (_params) {
                vc.component.refreshViewDbTableInfo();
                $('#viewDbTablesModel').modal('show');
                $that.viewDbTablesInfo.curDbId = _params.curDbId;
                $that.viewDbTablesInfo.curDbName = _params.curDbName;
                $that._loadTables();
            });
        },
        methods: {
            refreshViewDbTableInfo: function () {
                vc.component.viewDbTablesInfo = {
                    curDbId: '',
                    curDbName: '',
                    tables: []
                }
            },
            _loadTables: function () {
                let _data = {
                    dbId: $that.viewDbTablesInfo.curDbId,
                    sql: 'show tables'
                }
                vc.http.apiPost(
                    '/dbClient/execSql',
                    JSON.stringify(_data), {
                    emulateJSON: true
                },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        vc.toast(_json.msg);
                        if (_json.code == 0) {
                            _json.data.forEach(t => {
                                for (let key in t) {
                                    $that.viewDbTablesInfo.tables.push(t[key])
                                }
                            });
                        }
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            _selectTable:function(_tableName){
                vc.emit('mysqlClient','execQuerySql',{
                    dbId: $that.viewDbTablesInfo.curDbId,
                    curDbName: $that.viewDbTablesInfo.curDbName,
                    sql: 'select * from '+ _tableName +";"
                });
                $('#viewDbTablesModel').modal('hide');
            }
        }
    });

})(window.vc, window.vc.component);