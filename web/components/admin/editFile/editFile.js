(function (vc, vm) {

    vc.extends({
        data: {
            editFileInfo: {
                hostId: '',
                fileName: '',
                fileContext: '',
                curPath: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editFile', 'openEditFileModal', function (_params) {
                vc.component.refreshEditFileInfo();
                $('#editFileModel').modal('show');
                vc.copyObject(_params, vc.component.editFileInfo);
                $that._listFileContext();

            });
        },
        methods: {
            _listFileContext: function () {
                let _curPath = $that.editFileInfo.curPath;

                if (!_curPath.endsWith('/')) {
                    _curPath += ('/'+ $that.editFileInfo.fileName);
                }else{
                    _curPath += ( $that.editFileInfo.fileName);

                }
                let param ={
                    params:{
                        hostId:$that.editFileInfo.hostId,
                        fileName: _curPath
                    }
                }
                //发送get请求
                vc.http.apiGet('/host/listFileContext',
                    param,
                    function (json, res) {
                        let _businessPackageManageInfo = JSON.parse(json);
                        $that.editFileInfo.fileContext = _businessPackageManageInfo.data;
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            editFile: function () {
                let _curPath = $that.editFileInfo.curPath;
                if (!_curPath.endsWith('/')) {
                    _curPath += ('/' + $that.editFileInfo.fileName);
                }else{
                    _curPath += ( $that.editFileInfo.fileName);
                }
                let _data = {
                    hostId: $that.editFileInfo.hostId,
                    fileName: _curPath,
                    fileContext: $that.editFileInfo.fileContext,
                }
                vc.http.apiPost(
                    '/host/editFile',
                    JSON.stringify(_data), {
                    emulateJSON: true
                },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editFileModel').modal('hide');
                            vc.emit('fileManager', 'listFiles', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditFileInfo: function () {
                vc.component.editFileInfo = {
                    hostId: '',
                    fileName: '',
                    fileContext: '',
                    curPath: '',
                }
            },

        }
    });

})(window.vc, window.vc.component);