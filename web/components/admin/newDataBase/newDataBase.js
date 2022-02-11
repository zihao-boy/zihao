(function(vc) {

    vc.extends({
        data: {
            newDataBaseInfo: {
                name: '',
                ip: '%',
                username: '',
                password: '',
                curDbId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('newDataBase', 'openNewDataBaseModal', function(_param) {
                $that.newDataBaseInfo.curDbId = _param.curDbId;

                $('#newDataBaseModel').modal('show');
            });
        },
        methods: {
            newDataBaseValidate() {
                return vc.validate.validate({
                    newDataBaseInfo: vc.component.newDataBaseInfo
                }, {
                    'newDataBaseInfo.name': [{
                        limit: "required",
                        param: "",
                        errInfo: "数据库名不能为空"
                    }],
                    'newDataBaseInfo.ip': [{
                        limit: "required",
                        param: "",
                        errInfo: "访问主机不能为空"
                    }],
                    'newDataBaseInfo.username': [{
                            limit: "required",
                            param: "",
                            errInfo: "用户名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "用户名太长"
                        },
                    ],
                    'newDataBaseInfo.password': [{
                            limit: "required",
                            param: "",
                            errInfo: "密码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "密码太长"
                        },
                    ],
                });
            },
            _saveDataBase: function() {
                if (!vc.component.newDataBaseValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }


                let _newUser = "create user '" + $that.newDataBaseInfo.username + "'@'" + $that.newDataBaseInfo.ip + "' identified by '" + $that.newDataBaseInfo.password + "';"
                let _newDb = "CREATE DATABASE " + $that.newDataBaseInfo.name + " ;"
                let _userPri = "grant all privileges on " + $that.newDataBaseInfo.name + ".* to '" + $that.newDataBaseInfo.username + "'@'" + $that.newDataBaseInfo.ip + "' ;";
                let _flushPri = "flush privileges;"
                vc.emit('mysqlClient', 'execSql', {
                    dbId: $that.newDataBaseInfo.curDbId,
                    sql: _newUser + _newDb + _userPri + _flushPri
                })
                $('#newDataBaseModel').modal('hide');
            },
            clearNewDataBaseInfo: function() {
                vc.component.newDataBaseInfo = {
                    name: '',
                    ip: '%',
                    username: '',
                    password: ''
                };
            }
        }
    });

})(window.vc);