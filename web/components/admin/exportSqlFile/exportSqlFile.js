
(function (vc, vm) {

    vc.extends({
        data: {
            exportSqlFileInfo: {
                fileName: '',
                curDbId: '',
                error: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('exportSqlFile', 'openExportSqlModal', function (_params) {
                vc.component.clearExportSqlFileInfo();
                $('#exportSqlFileModel').modal('show');
                vc.copyObject(_params, vc.component.exportSqlFileInfo);
            });
        },
        methods: {
            _exportSqlFile: function () {
                $that.exportSqlFileInfo.error = ""

                let _fileName = vc.component.exportSqlFileInfo.fileName;

                if (!_fileName) {
                    vc.toast('文件名不能为空');
                    return;
                }
                let _data = {
                    dbId: $that.exportSqlFileInfo.curDbId,
                    fileName: _fileName
                }

                vc.http.apiPost(
                    '/dbClient/exportSqlFile',
                    JSON.stringify(_data), {
                    emulateJSON: true
                },
                    function (json, res) {
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            vc.toast(_json.data);
                            $('#exportSqlFileModel').modal('hide');
                            vc.emit('mysqlClient', 'showError', _json.data);
                            return;
                        }
                        vc.emit('mysqlClient', 'showError', _json.msg);
                        $that.exportSqlFileInfo.error = _json.msg
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                        //vc.toast(errInfo);
                        vc.emit('mysqlClient', 'showError', errInfo)
                        $that.exportSqlFileInfo.error = errInfo

                    });
            },
            clearExportSqlFileInfo: function () {
                vc.component.exportSqlFileInfo = {
                    fileName: '',
                    curDbId: '',
                    error: ''
                }
            },

        }
    });

})(window.vc, window.vc.component);