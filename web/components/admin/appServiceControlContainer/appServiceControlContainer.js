/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appServiceControlContainerInfo: {
                appServiceControlContainers: [],
                total: 0,
                records: 1,
                asId: '',
            }
        },
        _initMethod: function () {
        },
        _initEvent: function () {

            vc.on('appServiceControlContainer', 'switch', function (_param) {
                if (_param.asId == '') {
                    return;
                }
                vc.copyObject(_param, $that.appServiceControlContainerInfo)
                vc.component._listappServiceControlContainers(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('appServiceControlContainer', 'paginationPlus', 'page_event', function (_currentPage) {
                vc.component._listappServiceControlContainers(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listappServiceControlContainers: function (_page, _rows) {

                var param = {
                    params: {
                        page: _page,
                        row: _rows,
                        asId: $that.appServiceControlContainerInfo.asId
                    }
                };
                //发送get请求
                vc.http.apiGet('/appService/getAppServiceContainer',
                    param,
                    function (json, res) {
                        var _appServiceControlContainersInfo = JSON.parse(json);
                        vc.component.appServiceControlContainerInfo.total = _appServiceControlContainersInfo.total;
                        vc.component.appServiceControlContainerInfo.records = _appServiceControlContainersInfo.records;
                        vc.component.appServiceControlContainerInfo.appServiceControlContainers = _appServiceControlContainersInfo.data;
                        vc.emit('appServiceControlContainer', 'paginationPlus', 'init', {
                            total: vc.component.appServiceControlContainerInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }

        }
    });
})(window.vc);
