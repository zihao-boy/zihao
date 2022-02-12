
(function (vc, vm) {

    vc.extends({
        data: {
            importSqlFileInfo: {
                sqls: [],
                index: 0,
                curDbId: '',
                error:''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('importSqlFile', 'openImportSqlModal', function (_params) {
                vc.component.clearImportSqlFileInfo();
                $('#importSqlFileModel').modal('show');
                document.getElementById("files")[0].value = '';
                vc.copyObject(_params, vc.component.importSqlFileInfo);
            });
        },
        methods: {
            importSqlFile: function () {
                let selectedFile = document.getElementById('files').files[0];
                let reader = new FileReader(); //这是核心,读取操作就是由它完成.
                reader.readAsText(selectedFile); //读取文件的内容,也可以读取文件的URL
                reader.onload = function () {
                    //当读取完成后回调这个函数,然后此时文件的内容存储到了result中,直接操作即可
                    $that.importSqlFileInfo.sqls = this.result.split(";");
                    $that._doImportSqlFile();
                }
            },
            _doImportSqlFile: function () {
                $that.importSqlFileInfo.error = ""
                if ($that.importSqlFileInfo.index >= $that.importSqlFileInfo.sqls.length) {
                    //关闭model
                    $('#importSqlFileModel').modal('hide');
                    return;
                }

                let _sql = vc.component.importSqlFileInfo.sqls[$that.importSqlFileInfo.index];

                if (!_sql) {
                    $that.importSqlFileInfo.index += 1;
                    $that._doImportSqlFile();
                    return;
                }
                let _data = {
                    dbId: $that.importSqlFileInfo.curDbId,
                    sql: _sql
                }

                vc.http.apiPost(
                    '/dbClient/execSql',
                    JSON.stringify(_data), {
                    emulateJSON: true
                },
                    function (json, res) {
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            $that.importSqlFileInfo.index += 1;
                            $that._doImportSqlFile();
                            return;
                        }
                        vc.emit('mysqlClient', 'showError', _json.msg);
                        $that.importSqlFileInfo.error = _json.msg
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                        //vc.toast(errInfo);
                        vc.emit('mysqlClient', 'showError', errInfo)
                        $that.importSqlFileInfo.error = errInfo

                    });
            },
            clearImportSqlFileInfo: function () {
                vc.component.importSqlFileInfo = {
                    sqls: [],
                    index: 0,
                    curDbId: '',
                    error:''
                }
            },

        }
    });

})(window.vc, window.vc.component);