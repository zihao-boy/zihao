(function(vc) {

    vc.extends({
        data: {
            newTableInfo: {
                name: '',
                cols:[],
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
                if (!vc.component.newTableValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                let _newUser = "create user '"+$that.newTableInfo.username+"'@'"+$that.newTableInfo.ip+"' identified by '"+$that.newTableInfo.password+"';"
                let _newDb = "CREATE DATABASE "+$that.newTableInfo.name+" ;"
                let _userPri = "grant all privileges on "+$that.newTableInfo.name+".* to '"+$that.newTableInfo.username+"'@'"+$that.newTableInfo.ip+"' ;";
                let _flushPri = "flush privileges;"
                vc.emit('mysqlClient','execSql',{
                    dbId: $that.newTableInfo.curDbId,
                    sql: _newUser+_newDb+_userPri+_flushPri
                })
            },
            clearNewTableInfo: function() {
                vc.component.newTableInfo = {
                    name: '',
                    cols:[],
                };
            }
        }
    });

})(window.vc);