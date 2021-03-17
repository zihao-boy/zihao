(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostContainersInfo: {
                containers: [],
                hostId: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            //切换 至费用页面
            vc.on('hostContainers', 'switch', function (_param) {
                if (_param.hostId == '') {
                    return;
                }
                $that.clearhostContainersInfo();
                vc.copyObject(_param, $that.hostContainersInfo)
                $that._listhostContainers(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('hostContainers', 'listOwnerCarData', function (_param) {
                $that._listhostContainers(DEFAULT_PAGE, DEFAULT_ROWS);
            });
        },
        methods: {
            _listhostContainers: function (_page, _row) {

                let param = {
                    params: {
                        page: 1,
                        row: 50,
                        hostId: $that.hostContainersInfo.hostId
                    }
                }
                //发送get请求
                vc.http.apiGet('/host/getContainers',
                    param,
                    function (json, res) {
                        let _json = JSON.parse(json);
                        $that.hostContainersInfo.containers = _json.data;
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _addCarParkingSpace: function (_car) {
                vc.jumpToPage('/admin.html#/pages/property/carAddParkingSpace?carId=' + _car.carId);
            },
            clearhostContainersInfo: function () {
                $that.hostContainersInfo = {
                    containers: [],
                    hostId: ''
                }
            }

        }

    });
})(window.vc);
