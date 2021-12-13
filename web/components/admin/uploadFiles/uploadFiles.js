(function(vc, vm) {

    vc.extends({
        data: {
            uploadFilesInfo: {
                hostId: '',
                files: [],
                fileGroupName: 'd',
                curPath: '',
                index: 0
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('uploadFiles', 'openRenameFileModal', function(_params) {
                vc.component.refreshRenameFileInfo();
                $('#uploadFilesModel').modal('show');
                document.getElementsByName("files")[0].value = '';
                vc.copyObject(_params, vc.component.uploadFilesInfo);
            });
        },
        methods: {
            uploadFiles: function() {
                let _files = document.getElementsByName("files");
                $that.uploadFilesInfo.files = _files[0].files;
                $that._doUploadFile();
            },
            _doUploadFile: function() {
                if ($that.uploadFilesInfo.index >= $that.uploadFilesInfo.files.length) {
                    //关闭model
                    $('#uploadFilesModel').modal('hide');
                    vc.emit('fileManager', 'listFiles', {});
                    return;
                }
                let _curPath = $that.uploadFilesInfo.curPath;

                let _file = vc.component.uploadFilesInfo.files[$that.uploadFilesInfo.index];
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
                param.append('hostId', $that.uploadFilesInfo.hostId);
                param.append('curPath', _curPath);
                vc.http.apiPost(
                    '/host/uploadFile',
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
                            $that.uploadFilesInfo.index += 1;
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
                vc.component.uploadFilesInfo = {
                    hostId: '',
                    files: [],
                    fileGroupName: 'd',
                    curPath: '',
                    index: 0
                }
            },

        }
    });

})(window.vc, window.vc.component);