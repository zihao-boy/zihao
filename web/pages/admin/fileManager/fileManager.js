/**
    入驻小区
**/
(function(vc) {
    vc.extends({
        data: {
            fileManagerInfo: {
                curPath: '/',
                hostId: '',
                files: []
            }
        },
        _initMethod: function() {
            $that.fileManagerInfo.hostId = vc.getParam('hostId');

            $that._listFiles();

        },
        _initEvent: function() {
            vc.on('fileManager', 'listFiles', function() {
                $that._listFiles();
            });

        },
        methods: {
            _changeDir: function(_file) {
                if (_file.groupName != 'd') {
                    $that._openEditFile(_file); //修改文件
                    return;
                }
                if ($that.fileManagerInfo.curPath.endsWith('/')) {
                    $that.fileManagerInfo.curPath += _file.name;
                } else {
                    $that.fileManagerInfo.curPath += ("/" + _file.name);
                }
                $that._listFiles();
            },
            _parentDir: function() {
                //如果是跟目录 直接显示根目录
                if (!$that.fileManagerInfo.curPath || $that.fileManagerInfo.curPath == '/') {
                    $that.fileManagerInfo.curPath = "/";
                    $that._listFiles();
                    return;
                }

                if ($that.fileManagerInfo.curPath.endsWith('/')) {
                    $that.fileManagerInfo.curPath = $that.fileManagerInfo.curPath.substring(0, $that.fileManagerInfo.curPath.length - 1);
                }
                $that.fileManagerInfo.curPath = $that.fileManagerInfo.curPath.substring(0, $that.fileManagerInfo.curPath.lastIndexOf('/'));

                //如果返回到最后还是根目录 那还是 显示根目录
                if (!$that.fileManagerInfo.curPath) {
                    $that.fileManagerInfo.curPath = "/";
                }
                $that._listFiles();
            },

            _listFiles: function() {
                if (!$that.fileManagerInfo.curPath) {
                    return;
                }
                let param = {
                    params: {
                        hostId: $that.fileManagerInfo.hostId,
                        curPath: $that.fileManagerInfo.curPath
                    }
                }

                //发送get请求
                vc.http.apiGet('/host/listFiles',
                    param,
                    function(json, res) {
                        let _businessPackageManageInfo = JSON.parse(json);
                        $that.fileManagerInfo.files = _businessPackageManageInfo.data;
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _deleteFileOrDir: function(_file) {
                let _curPath = $that.fileManagerInfo.curPath;
                if (!_curPath.endsWith('/')) {
                    _curPath += ('/' + _file.name);
                }
                let _data = {
                    hostId: $that.fileManagerInfo.hostId,
                    fileName: _curPath,
                    fileGroupName: _file.groupName
                }
                vc.emit('deleteFileOrDir', 'openDeleteFileOrDirModal', _data);
            },
            openNewFile: function(isFile) {
                let _curPath = $that.fileManagerInfo.curPath;
                vc.emit('newFile', 'openNewFileModal', {
                    hostId: $that.fileManagerInfo.hostId,
                    fileGroupName: isFile,
                    curPath: _curPath
                })
            },

            _openRenameFileModel: function(_file) {
                let _curPath = $that.fileManagerInfo.curPath;
                vc.emit('renameFile', 'openRenameFileModal', {
                    hostId: $that.fileManagerInfo.hostId,
                    fileGroupName: _file.groupName,
                    fileName: _file.name,
                    curPath: _curPath
                })
            },
            _openEditFile: function(_file) {
                let _curPath = $that.fileManagerInfo.curPath;
                if (_file.size > 1024 * 1024) {
                    vc.toast('文件超过1M,不能在线修改，请下载修改');
                    return;
                }
                vc.emit('editFile', 'openEditFileModal', {
                    hostId: $that.fileManagerInfo.hostId,
                    fileName: _file.name,
                    curPath: _curPath
                })
            },
            _openFiles: function(isFile) {
                let _curPath = $that.fileManagerInfo.curPath;
                vc.emit('uploadFiles', 'openRenameFileModal', {
                    hostId: $that.fileManagerInfo.hostId,
                    fileGroupName: isFile,
                    curPath: _curPath
                });
            }
        }
    });
})(window.vc);