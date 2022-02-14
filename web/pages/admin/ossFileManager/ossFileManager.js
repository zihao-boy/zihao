/**
    入驻小区
**/
(function(vc) {
    vc.extends({
        data: {
            ossFileManagerInfo: {
                curPath: '/',
                ossId: '',
                files: [],
                ossName: ''
            }
        },
        _initMethod: function() {
            $that.ossFileManagerInfo.ossId = vc.getParam('ossId');
            $that.ossFileManagerInfo.ossName = vc.getParam('ossName');
            let _cacheCurPath = vc.getData($that.ossFileManagerInfo.ossId + '_curPath');
            if (_cacheCurPath) {
                $that.ossFileManagerInfo.curPath = _cacheCurPath;
            }
            $that._listFiles();

        },
        _initEvent: function() {
            vc.on('ossFileManager', 'listFiles', function() {
                $that._listFiles();
            });

        },
        methods: {
            _changeDir: function(_file) {
                if (_file.groupName != 'd') {
                    return;
                }
                if ($that.ossFileManagerInfo.curPath.endsWith('/')) {
                    $that.ossFileManagerInfo.curPath += _file.name;
                } else {
                    $that.ossFileManagerInfo.curPath += ("/" + _file.name);
                }
                $that._listFiles();
            },
            _parentDir: function() {
                //如果是跟目录 直接显示根目录
                if (!$that.ossFileManagerInfo.curPath || $that.ossFileManagerInfo.curPath == '/') {
                    $that.ossFileManagerInfo.curPath = "/";
                    $that._listFiles();
                    return;
                }

                if ($that.ossFileManagerInfo.curPath.endsWith('/')) {
                    $that.ossFileManagerInfo.curPath = $that.ossFileManagerInfo.curPath.substring(0, $that.ossFileManagerInfo.curPath.length - 1);
                }
                $that.ossFileManagerInfo.curPath = $that.ossFileManagerInfo.curPath.substring(0, $that.ossFileManagerInfo.curPath.lastIndexOf('/'));

                if (!$that.ossFileManagerInfo.curPath.endsWith("/")) {
                    $that.ossFileManagerInfo.curPath += "/";
                }

                //如果返回到最后还是根目录 那还是 显示根目录
                if (!$that.ossFileManagerInfo.curPath) {
                    $that.ossFileManagerInfo.curPath = "/";
                }
                $that._listFiles();
            },

            _listFiles: function() {
                if (!$that.ossFileManagerInfo.curPath) {
                    return;
                }
                //存起来方便操作
                vc.saveData($that.ossFileManagerInfo.ossId + '_curPath', $that.ossFileManagerInfo.curPath);
                let param = {
                    params: {
                        ossId: $that.ossFileManagerInfo.ossId,
                        curPath: $that.ossFileManagerInfo.curPath
                    }
                }

                //发送get请求
                vc.http.apiGet('/resources/listOssFiles',
                    param,
                    function(json, res) {
                        let _businessPackageManageInfo = JSON.parse(json);
                        $that.ossFileManagerInfo.files = _businessPackageManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _deleteFileOrDir: function(_file) {
                let _curPath = $that.ossFileManagerInfo.curPath;
                if (!_curPath.endsWith('/')) {
                    _curPath += ('/' + _file.name);
                } else {
                    _curPath += (_file.name);
                }
                let _data = {
                    ossId: $that.ossFileManagerInfo.ossId,
                    curPath: _curPath,
                    fileGroupName: _file.groupName
                }
                vc.emit('deleteFileOrDir', 'openDeleteFileOrDirModal', _data);
            },
            openNewFile: function(isFile) {
                let _curPath = $that.ossFileManagerInfo.curPath;
                vc.emit('newFile', 'openNewFileModal', {
                    ossId: $that.ossFileManagerInfo.ossId,
                    fileGroupName: isFile,
                    curPath: _curPath
                })
            },

            _openRenameFileModel: function(_file) {
                let _curPath = $that.ossFileManagerInfo.curPath;
                vc.emit('renameFile', 'openRenameFileModal', {
                    ossId: $that.ossFileManagerInfo.ossId,
                    fileGroupName: _file.groupName,
                    fileName: _file.name,
                    curPath: _curPath
                })
            },
            _openFiles: function(isFile) {
                let _curPath = $that.ossFileManagerInfo.curPath;
                vc.emit('uploadOssFiles', 'openUploadOssFilesModal', {
                    ossId: $that.ossFileManagerInfo.ossId,
                    fileGroupName: isFile,
                    curPath: _curPath
                });
            },
            _downLoadFile: function(file) {
                vc.jumpToPage('/app/resources/downloadOssFile?ossId=' + $that.ossFileManagerInfo.ossId + "&curPath=" + $that.ossFileManagerInfo.curPath + "&fileName=" + file.name);
            },
            _goBack: function() {
                vc.goBack();
            },
        }
    });
})(window.vc);