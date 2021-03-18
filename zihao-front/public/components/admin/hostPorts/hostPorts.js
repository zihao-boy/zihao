(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostPortsInfo: {
                ports: [],
                hostId: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            //切换 至费用页面
            vc.on('hostPorts', 'switch', function (_param) {
                if (_param.hostId == '') {
                    return;
                }
                $that.clearhostPortsInfo();
                vc.copyObject(_param, $that.hostPortsInfo)
                $that._listhostPorts(DEFAULT_PAGE, DEFAULT_ROWS);
            });
        },
        methods: {
            _listhostPorts: function (_page, _row) {

                let param = {
                    params: {
                        page: 1,
                        row: 50,
                        hostId: $that.hostPortsInfo.hostId
                    }
                }
                //发送get请求
                vc.http.apiGet('/host/getHostPort',
                    param,
                    function (json, res) {
                        let _json = JSON.parse(json);
                        $that.hostPortsInfo.ports = _json.data;
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            clearhostPortsInfo: function () {
                $that.hostPortsInfo = {
                    ports: [],
                    hostId: ''
                }
            },
            

        }

    });
})(window.vc);
