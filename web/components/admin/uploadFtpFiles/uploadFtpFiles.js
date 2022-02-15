(function(vc, vm) {

    vc.extends({
        data: {
            uploadFtpFilesInfo: {
                ftpId: '',
                files: [],
                fileGroupName: 'd',
                curPath: '',
                index: 0
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('uploadFtpFiles', 'openUploadFtpFilesModal', function(_params) {
                vc.component.refreshRenameFileInfo();
                $('#uploadFtpFilesModel').modal('show');
                document.getElementsByName("files")[0].value = '';
                vc.copyObject(_params, vc.component.uploadFtpFilesInfo);
            });
        },
        methods: {
            uploadFtpFiles: function() {
                let _files = document.getElementsByName("files");
                $that.uploadFtpFilesInfo.files = _files[0].files;
                $that._doUploadFile();
            },
            _doUploadFile: function() {
                if ($that.uploadFtpFilesInfo.index >= $that.uploadFtpFilesInfo.files.length) {
                    //关闭model
                    $('#uploadFtpFilesModel').modal('hide');
                    vc.emit('ftpFileManager', 'listFiles', {});
                    return;
                }
                let _curPath = $that.uploadFtpFilesInfo.curPath;

                let _file = vc.component.uploadFtpFilesInfo.files[$that.uploadFtpFilesInfo.index];
                console.log(_file)

                let _relativePath = _file.webkitRelativePath;

                _relativePath = _relativePath.substring(0, _relativePath.lastIndexOf('/'))

                if (!_curPath.endsWith('/')) {
                    _curPath += ('/' + _relativePath);
                } else {
                    _curPath += (_relativePath);
                }
                let param = new FormData();
                param.append("uploadFile", _file);
                param.append('ftpId', $that.uploadFtpFilesInfo.ftpId);
                param.append('curPath', _curPath);
                vc.http.apiPost(
                    '/resources/uploadFtpFile',
                    param, {
                        emulateJSON: true,
                        //添加请求头
                        headers: {
                            "Content-Type": "multipart/form-data"
                        }
                    },
                    function(json, res) {
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            $that.uploadFtpFilesInfo.index += 1;
                            $that._doUploadFile();
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            refreshRenameFileInfo: function() {
                vc.component.uploadFtpFilesInfo = {
                    ftpId: '',
                    files: [],
                    fileGroupName: 'd',
                    curPath: '',
                    index: 0
                }
            },

        }
    });

})(window.vc, window.vc.component);