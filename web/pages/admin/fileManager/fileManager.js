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