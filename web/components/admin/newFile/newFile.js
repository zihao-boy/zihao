(function(vc, vm) {

    vc.extends({
        data: {
            newFileInfo: {
                hostId:'',
                fileName:'',
                fileGroupName:'',
                curPath:''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('newFile', 'openNewFileModal', function(_params) {
                vc.component.refreshNewFileInfo();
                $('#newFileModel').modal('show');
                vc.copyObject(_params, vc.component.newFileInfo);
            });
        },
        methods: {
            newFile: function() {
                let _curPath = $that.newFileInfo.curPath;
                if (!_curPath.endsWith('/')) {
                    _curPath += ('/'+ $that.newFileInfo.fileName);
                }
                let _data = {
                    hostId:$that.newFileInfo.hostId,
                    fileName:_curPath,
                    fileGroupName:$that.newFileInfo.fileGroupName,
                }
                vc.http.apiPost(
                    '/host/newFile',
                    JSON.stringify(_data), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#newFileModel').modal('hide');
                            vc.emit('fileManager','listFiles', {});
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
                    hostId:'',
                    fileName:'',
                    fileGroupName:'',
                    curPath:'',
                }
            },
           
        }
    });

})(window.vc, window.vc.component);