(function(vc) {

    vc.extends({
        data: {
            newTableInfo: {
                name: '',
                cols:[],
                sql:'',
                view:'form'
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('newTable', 'openNewTableModal', function(_param) {
                $that.newTableInfo.curDbId = _param.curDbId;
                $('#newTableModel').modal('show');
            });
        },
        methods: {
            _saveDataBase: function() {
                
                $that._doGetCreateTableSql();
                vc.emit('mysqlClient','execSql',{
                    dbId: $that.newTableInfo.curDbId,
                    sql: $that.newTableInfo.sql
                })
            },
            _doGetCreateTableSql:function(){
                let _newUser = "create table "+$that.newTableInfo.name+"( \n";
                $that.newTableInfo.cols.forEach(item => {
                    _newUser += (item.name + " " + item.type);
                    if(item.type == 'bigint' || item.type == 'int'){
                        //不写
                    }else{
                        _newUser += ("("+item.len+")")
                    }

                    if(item.isNull == 'N'){
                        _newUser += " not null "
                    }

                    if(item.isPrimay == 'Y'){
                        _newUser += " PRIMARY KEY "
                    }

                    if(item.default){
                        _newUser += " default " + item.default
                    }

                    if(item.comment){
                        _newUser += " comment '" + item.comment+"'"
                    }
                    _newUser += ",\n"
                });
                if(_newUser.endsWith(',\n')){
                    _newUser = _newUser.substring(0,_newUser.length-2) + "\n";
                }
                _newUser += ");"
                $that.newTableInfo.sql = _newUser;
            },
            clearNewTableInfo: function() {
                vc.component.newTableInfo = {
                    name: '',
                    cols:[],
                };
            },
            _addCol:function(){
                $that.newTableInfo.cols.push({
                    id: vc.uuid(),
                    name:'',
                    type:'',
                    len:0,
                    isNull:'N',
                    isPrimay:'N',
                    default:'',
                    comment:''
                })
            },
            switchSqlView:function(_view){
                $that.newTableInfo.view = _view;
                $that._doGetCreateTableSql();
            }
        }
    });

})(window.vc);