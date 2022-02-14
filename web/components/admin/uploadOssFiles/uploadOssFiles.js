(function(vc, vm) {

    vc.extends({
        data: {
            uploadOssFilesInfo: {
                ossId: '',
                files: [],
                fileGroupName: 'd',
                curPath: '',
                index: 0
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('uploadOssFiles', 'openUploadOssFilesModal', function(_params) {
                vc.component.refreshRenameFileInfo();
                $('#uploadOssFilesModel').modal('show');
                document.getElementsByName("files")[0].value = '';
                vc.copyObject(_params, vc.component.uploadOssFilesInfo);
            });
        },
        methods: {
            uploadOssFiles: function() {
                let _files = document.getElementsByName("files");
                $that.uploadOssFilesInfo.files = _files[0].files;
                $that._doUploadFile();
            },
            _doUploadFile: function() {
                if ($that.uploadOssFilesInfo.index >= $that.uploadOssFilesInfo.files.length) {
                    //关闭model
                    $('#uploadOssFilesModel').modal('hide');
                    vc.emit('ossFileManager', 'listFiles', {});
                    return;
                }
                let _curPath = $that.uploadOssFilesInfo.curPath;

                let _file = vc.component.uploadOssFilesInfo.files[$that.uploadOssFilesInfo.index];
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
                param.append('ossId', $that.uploadOssFilesInfo.ossId);
                param.append('curPath', _curPath);
                vc.http.apiPost(
                    '/resources/uploadOssFile',
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
                            $that.uploadOssFilesInfo.index += 1;
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
                vc.component.uploadOssFilesInfo = {
                    ossId: '',
                    files: [],
                    fileGroupName: 'd',
                    curPath: '',
                    index: 0
                }
            },

        }
    });

})(window.vc, window.vc.component);