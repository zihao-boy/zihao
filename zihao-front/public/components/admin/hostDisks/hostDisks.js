(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostDisksInfo: {
                disks: [],
                hostId: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            //切换 至费用页面
            vc.on('hostDisks', 'switch', function (_param) {
                if (_param.hostId == '') {
                    return;
                }
                $that.clearhostDisksInfo();
                vc.copyObject(_param, $that.hostDisksInfo)
                $that._listhostDisks(DEFAULT_PAGE, DEFAULT_ROWS);
            });
        },
        methods: {
            _listhostDisks: function (_page, _row) {

                let param = {
                    params: {
                        page: 1,
                        row: 50,
                        hostId: $that.hostDisksInfo.hostId
                    }
                }
                //发送get请求
                vc.http.apiGet('/host/getHostDisk',
                    param,
                    function (json, res) {
                        let _json = JSON.parse(json);
                        $that.hostDisksInfo.disks = _json.data;
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            clearhostDisksInfo: function () {
                $that.hostDisksInfo = {
                    disks: [],
                    hostId: ''
                }
            },
            

        }

    });
})(window.vc);
