/**
    入驻小区
**/
(function(vc) {
    vc.extends({
        data: {
            ftpFileManagerInfo: {
                curPath: '/',
                ftpId: '',
                files: [],
                ftpName: ''
            }
        },
        _initMethod: function() {
            $that.ftpFileManagerInfo.ftpId = vc.getParam('ftpId');
            $that.ftpFileManagerInfo.ftpName = vc.getParam('ftpName');
            let _cacheCurPath = vc.getData($that.ftpFileManagerInfo.ftpId + '_curPath');
            if (_cacheCurPath) {
                $that.ftpFileManagerInfo.curPath = _cacheCurPath;
            }
            $that._listFiles();

        },
        _initEvent: function() {
            vc.on('ftpFileManager', 'listFiles', function() {
                $that._listFiles();
            });

        },
        methods: {
            _changeDir: function(_file) {
                if (_file.groupName != 'd') {
                    return;
                }
                if ($that.ftpFileManagerInfo.curPath.endsWith('/')) {
                    $that.ftpFileManagerInfo.curPath += _file.name;
                } else {
                    $that.ftpFileManagerInfo.curPath += ("/" + _file.name);
                }
                $that._listFiles();
            },
            _parentDir: function() {
                //如果是跟目录 直接显示根目录
                if (!$that.ftpFileManagerInfo.curPath || $that.ftpFileManagerInfo.curPath == '/') {
                    $that.ftpFileManagerInfo.curPath = "/";
                    $that._listFiles();
                    return;
                }

                if ($that.ftpFileManagerInfo.curPath.endsWith('/')) {
                    $that.ftpFileManagerInfo.curPath = $that.ftpFileManagerInfo.curPath.substring(0, $that.ftpFileManagerInfo.curPath.length - 1);
                }
                $that.ftpFileManagerInfo.curPath = $that.ftpFileManagerInfo.curPath.substring(0, $that.ftpFileManagerInfo.curPath.lastIndexOf('/'));

                if (!$that.ftpFileManagerInfo.curPath.endsWith("/")) {
                    $that.ftpFileManagerInfo.curPath += "/";
                }

                //如果返回到最后还是根目录 那还是 显示根目录
                if (!$that.ftpFileManagerInfo.curPath) {
                    $that.ftpFileManagerInfo.curPath = "/";
                }
                $that._listFiles();
            },

            _listFiles: function() {
                if (!$that.ftpFileManagerInfo.curPath) {
                    return;
                }
                //存起来方便操作
                vc.saveData($that.ftpFileManagerInfo.ftpId + '_curPath', $that.ftpFileManagerInfo.curPath);
                let param = {
                    params: {
                        ftpId: $that.ftpFileManagerInfo.ftpId,
                        curPath: $that.ftpFileManagerInfo.curPath
                    }
                }

                //发送get请求
                vc.http.apiGet('/resources/listFtpFiles',
                    param,
                    function(json, res) {
                        let _businessPackageManageInfo = JSON.parse(json);
                        $that.ftpFileManagerInfo.files = _businessPackageManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _deleteFileOrDir: function(_file) {
                let _curPath = $that.ftpFileManagerInfo.curPath;
                if (!_curPath.endsWith('/')) {
                    _curPath += ('/' + _file.name);
                } else {
                    _curPath += (_file.name);
                }
                let _data = {
                    ftpId: $that.ftpFileManagerInfo.ftpId,
                    curPath: _curPath,
                    fileGroupName: _file.groupName
                }
                vc.emit('deleteFileOrDir', 'openDeleteFileOrDirModal', _data);
            },
            openNewFile: function(isFile) {
                let _curPath = $that.ftpFileManagerInfo.curPath;
                vc.emit('newFtpFile', 'openNewFileModal', {
                    ftpId: $that.ftpFileManagerInfo.ftpId,
                    fileGroupName: isFile,
                    curPath: _curPath
                })
            },

            _openRenameFileModel: function(_file) {
                let _curPath = $that.ftpFileManagerInfo.curPath;
                vc.emit('renameFtpFile', 'openRenameFileModal', {
                    ftpId: $that.ftpFileManagerInfo.ftpId,
                    fileGroupName: _file.groupName,
                    fileName: _file.name,
                    curPath: _curPath
                })
            },
            _openFiles: function(isFile) {
                let _curPath = $that.ftpFileManagerInfo.curPath;
                vc.emit('uploadFtpFiles', 'openUploadFtpFilesModal', {
                    ftpId: $that.ftpFileManagerInfo.ftpId,
                    fileGroupName: isFile,
                    curPath: _curPath
                });
            },
            _downLoadFile: function(file) {
                vc.jumpToPage('/app/resources/downloadFtpFile?ftpId=' + $that.ftpFileManagerInfo.ftpId + "&curPath=" + $that.ftpFileManagerInfo.curPath + "&fileName=" + file.name);
            },
            _goBack: function() {
                vc.goBack();
            },
        }
    });
})(window.vc);