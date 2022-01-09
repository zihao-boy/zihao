(function(vc) {

    vc.extends({
        data: {
            mysqlClientInfo: {
                dbLinks: [],
                curDbId: '',
                curDbName: '无',
                _currentTab: {},
                _currentSql: '',
                sqlTabs: [],
                dataCols:[],
                data:[]

            }
        },
        _initMethod: function() {

            $("#text").setTextareaCount({
                width: "30px",
                bgColor: "#FFF",
                color: "#000",
                display: "inline-block"
            });

            $that._loadDbLink();
        },
        _initEvent: function() {

            vc.on('mysqlClient', 'load', function() {
                $that._loadDbLink();
            })

            vc.on('mysqlClient','execSql',function(_data){
                $that._doExecSql(_data);
            })
            vc.on('mysqlClient','execQuerySql',function(_data){
                $that._chooseDb({
                    id:_data.dbId,
                    name:_data.curDbName,
                    sqlText:_data.sql
                });
                $that._doExecSql(_data);
                
            })

        },
        watch: {
            'mysqlClientInfo._currentSql': {
                deep: true,
                handler: function(_new, _old) {
                    // console.log(_new);
                    $that.mysqlClientInfo._currentTab.sqlText = _new;
                }
            }
        },
        methods: {
            changeTab: function(_tab) {
                $that.mysqlClientInfo._currentTab = _tab;
                $that.mysqlClientInfo._currentSql = _tab.sqlText;
                $that.mysqlClientInfo.curDbId = _tab.curDbId;
                $that.mysqlClientInfo.curDbName = _tab.name;
            },
            _deleteTab:function(_tab,_index){
                $that.mysqlClientInfo.sqlTabs.splice(_index,1);
                if($that.mysqlClientInfo.sqlTabs.length < 1){
                    $that.mysqlClientInfo.curDbId = '';
                    $that.mysqlClientInfo.curDbName = '无';
                    $that.mysqlClientInfo._currentTab={};
                    return;
                }
                $that.changeTab($that.mysqlClientInfo.sqlTabs[0])
            },
            _customKeypress: function() {
                let typeSql = window.getSelection().toString();
                if (!$that.mysqlClientInfo.curDbId) {
                    vc.toast('请先选择数据库');
                    return;
                }
                if(!typeSql){
                    typeSql = $that.mysqlClientInfo._currentTab.sqlText;
                }
                $that.mysqlClientInfo.data = [];
                let _data = {
                    dbId: $that.mysqlClientInfo.curDbId,
                    sql: typeSql
                }
                $that._doExecSql(_data);
            },
            _doExecSql:function(_data){
                vc.http.apiPost(
                    '/dbClient/execSql',
                    JSON.stringify(_data), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        vc.toast(_json.msg);
                        if (_json.code == 0) {
                            $that._dealSqlData(_json.data);
                        }
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            _dealSqlData:function(data){
                //关闭model
                if(data.length <1){
                    return ;
                }
                $that.mysqlClientInfo.dataCols = Object.keys(data[0]);
                $that.mysqlClientInfo.data = data;
            },
            _loadDbLink: function() {
                let _param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };
                //发送get请求
                vc.http.apiGet('/dbClient/getDbLink',
                    _param,
                    function(json, res) {
                        var _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.mysqlClientInfo.dbLinks = _appVarGroupManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openViewDbDataModel:function(data){
                vc.emit('viewDbData', 'openViewDbDataModal',{
                    dataCols:$that.mysqlClientInfo.dataCols,
                    data:data
                })
            },
            _openNewDbLinkModal: function() {
                vc.emit('newDbLink', 'openNewDbLinkModal', {})
            },
            _openEditDbLinkModal: function(_dbLink) {
                vc.emit('editDbLink', 'openEditDbLinkModal', _dbLink);
            },
            _openViewDbTablesModal:function(_dbLink){
                vc.emit('viewDbTables', 'openViewDbDataModal', {
                    curDbId:_dbLink.id,
                    curDbName: _dbLink.name
                });
            },
            _openDeleteDbLinkModal: function(_dbLink) {
                vc.emit('deleteDbLink', 'openDeleteDbLinkModal', _dbLink);
            },
            _chooseDb: function(_dbLink) {
                $that.mysqlClientInfo.curDbId = _dbLink.id;
                $that.mysqlClientInfo.curDbName = _dbLink.name;
                let _id = vc.uuid();
                let _sqlText = '';
                if(_dbLink.sqlText){
                    _sqlText = _dbLink.sqlText;
                }
                $that.mysqlClientInfo._currentTab = {
                    id: _id,
                    name: _dbLink.name,
                    dbId: _dbLink.id,
                    sqlText: _sqlText,
                    curDbId: _dbLink.id,
                    curDbName: _dbLink.name,
                };
                $that.mysqlClientInfo.sqlTabs.push($that.mysqlClientInfo._currentTab);
                //$that.mysqlClientInfo._currentTab = _id;
                $that.mysqlClientInfo._currentSql = _sqlText;
            },
            _newDbQuery: function() {
                if (!$that.mysqlClientInfo.curDbId) {
                    vc.toast('请先选择数据库');
                    return;
                }
                let _id = vc.uuid();
                let _sqlText = '';
                $that.mysqlClientInfo._currentTab = {
                    id: _id,
                    name: $that.mysqlClientInfo.curDbName,
                    dbId: $that.mysqlClientInfo.curDbId,
                    sqlText: _sqlText,
                    curDbId: $that.mysqlClientInfo.curDbId,
                    curDbName: $that.mysqlClientInfo.curDbName,
                };
                $that.mysqlClientInfo.sqlTabs.push($that.mysqlClientInfo._currentTab);
                //$that.mysqlClientInfo._currentTab = _id;
                $that.mysqlClientInfo._currentSql = _sqlText;
            },
            _newDataBase:function(){
                if (!$that.mysqlClientInfo.curDbId) {
                    vc.toast('请先选择数据库');
                    return;
                }

                vc.emit('newDataBase', 'openNewDataBaseModal',{
                    curDbId:$that.mysqlClientInfo.curDbId,
                    curDbName:$that.mysqlClientInfo.curDbId,
                })

            }

        }
    });

})(window.vc);