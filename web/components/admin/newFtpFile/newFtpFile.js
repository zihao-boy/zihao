(function(vc, vm) {

    vc.extends({
        data: {
            newFileInfo: {
                ftpId: '',
                fileName: '',
                fileGroupName: '',
                curPath: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('newFtpFile', 'openNewFileModal', function(_params) {
                vc.component.refreshNewFileInfo();
                $('#newFileModel').modal('show');
                vc.copyObject(_params, vc.component.newFileInfo);
            });
        },
        methods: {
            newFile: function() {
                let _curPath = $that.newFileInfo.curPath;
                if (!_curPath.endsWith('/')) {
                    _curPath += ('/' + $that.newFileInfo.fileName);
                } else {
                    _curPath += ($that.newFileInfo.fileName);
                }
                let _data = {
                    ftpId: $that.newFileInfo.ftpId,
                    curPath: _curPath,
                    fileGroupName: $that.newFileInfo.fileGroupName,
                }
                vc.http.apiPost(
                    '/resources/newFtpFile',
                    JSON.stringify(_data), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#newFileModel').modal('hide');
                            vc.emit('ftpFileManager', 'listFiles', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshNewFileInfo: function() {
                vc.component.newFileInfo = {
                    ftpId: '',
                    fileName: '',
                    fileGroupName: '',
                    curPath: '',
                }
            },

        }
    });

})(window.vc, window.vc.component);