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

        },
        methods: {
            _changeDir: function(_file) {
                if (_file.groupName != 'd') {
                    return;
                }
                if ($that.fileManagerInfo.curPath.endsWith('/')) {
                    $that.fileManagerInfo.curPath += _file.name;
                } else {
                    $that.fileManagerInfo += ("/" + _file.name);
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
            }
        }
    });
})(window.vc);