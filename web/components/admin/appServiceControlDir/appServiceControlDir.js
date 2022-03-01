/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appServiceControlDirInfo: {
                appServiceControlDirs: [],
                total: 0,
                records: 1,
                asId: '',
                hostId: '',
                groupId: '',
                curSrcDir: '',
                flag: ''
            }
        },
        _initMethod: function() {},
        _initEvent: function() {

            vc.on('appServiceControlDir', 'switch', function(_param) {
                if (_param.asId == '') {
                    return;
                }
                vc.copyObject(_param, $that.appServiceControlDirInfo);
                if (_param.hasOwnProperty('asDeployType') && '1001' == _param.asDeployType) {
                    $that.appServiceControlDirInfo.groupId = _param.asDeployId;
                } else {
                    $that.appServiceControlDirInfo.hostId = _param.asDeployId;
                }
                vc.component._listappServiceControlDirs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('appServiceControlDir', 'paginationPlus', 'page_event', function(_currentPage) {
                vc.component._listappServiceControlDirs(_currentPage, DEFAULT_ROWS);
            });

            vc.on('appServiceControlDir', 'chooseHost', function(_host) {
                if ($that.appServiceControlDirInfo.flag == "toEditFile") {
                    setTimeout(function() {
                        $that._editFileContext(_host);
                    }, 1000); //参数是字符串

                    return;
                }
                $that._jumpToHostDir(_host);
            })
        },
        methods: {
            _listappServiceControlDirs: function(_page, _rows) {

                var param = {
                    params: {
                        page: _page,
                        row: _rows,
                        asId: $that.appServiceControlDirInfo.asId,

                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppServiceDir',
                    param,
                    function(json, res) {
                        var _appServiceControlDirsInfo = JSON.parse(json);
                        vc.component.appServiceControlDirInfo.total = _appServiceControlDirsInfo.total;
                        vc.component.appServiceControlDirInfo.records = _appServiceControlDirsInfo.records;
                        vc.component.appServiceControlDirInfo.appServiceControlDirs = _appServiceControlDirsInfo.data;
                        vc.emit('appServiceControlDir', 'paginationPlus', 'init', {
                            total: vc.component.appServiceControlDirInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddDirModal: function() {
                vc.emit('addAppServiceDir', 'openAddAppServiceDirModal', $that.appServiceControlDirInfo);
            },
            _openUpdateDirModal: function(event) {
                vc.emit('editAppServiceDir', 'openEditAppServiceDirModal', event);
            },
            _openDeleteDirModal: function(_dir) {
                vc.emit('deleteAppServiceDir', 'openDeleteAppServiceDirModal', _dir);
            },
            _toQuicklyDir: function(_dir) {
                $that.appServiceControlDirInfo.curSrcDir = _dir.srcDir;
                $that.appServiceControlDirInfo.flag = "toQuicklyDir";
                vc.emit('chooseHost', 'openChooseHostModel', {
                    hostId: $that.appServiceControlDirInfo.hostId,
                    groupId: $that.appServiceControlDirInfo.groupId,
                })
            },
            _toEditFile: function(_dir) {
                $that.appServiceControlDirInfo.curSrcDir = _dir.srcDir;
                $that.appServiceControlDirInfo.flag = "toEditFile";

                vc.emit('chooseHost', 'openChooseHostModel', {
                    hostId: $that.appServiceControlDirInfo.hostId,
                    groupId: $that.appServiceControlDirInfo.groupId,
                })
            },
            _jumpToHostDir: function(_host) {
                // 判断是文件夹还是文件
                vc.saveData(_host.hostId + "_curPath", $that.appServiceControlDirInfo.curSrcDir);
                vc.jumpToPage('/index.html#/pages/admin/fileManager?hostId=' + _host.hostId);
            },
            _editFileContext: function(_host) {
                let _curSrcDir = $that.appServiceControlDirInfo.curSrcDir;
                let _fileName = '';
                let _curPath = '';
                if (_curSrcDir.lastIndexOf('/') > 0) {
                    _curPath = _curSrcDir.substring(1, _curSrcDir.lastIndexOf('/'))
                    _fileName = _curSrcDir.substring(_curSrcDir.lastIndexOf('/') + 1, _curSrcDir.length)
                } else {
                    _curPath = "/"
                    _fileName = _curSrcDir.substring(1, _curSrcDir.length)
                }
                vc.emit('editFile', 'openEditFileModal', {
                    hostId: _host.hostId,
                    fileName: _fileName,
                    curPath: _curPath
                })
            }

        }
    });
})(window.vc);