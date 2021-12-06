/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appServiceControlHostsInfo: {
                appServiceControlHostss: [],
                total: 0,
                records: 1,
                asId: '',
            }
        },
        _initMethod: function () {
        },
        _initEvent: function () {

            vc.on('appServiceControlHosts', 'switch', function (_param) {
                if (_param.asId == '') {
                    return;
                }
                vc.copyObject(_param, $that.appServiceControlHostsInfo)
                vc.component._listappServiceControlHostss(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('appServiceControlHosts', 'paginationPlus', 'page_event', function (_currentPage) {
                vc.component._listappServiceControlHostss(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listappServiceControlHostss: function (_page, _rows) {

                var param = {
                    params: {
                        page: _page,
                        row: _rows,
                        asId: $that.appServiceControlHostsInfo.asId
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppServiceHosts',
                    param,
                    function (json, res) {
                        var _appServiceControlHostssInfo = JSON.parse(json);
                        vc.component.appServiceControlHostsInfo.total = _appServiceControlHostssInfo.total;
                        vc.component.appServiceControlHostsInfo.records = _appServiceControlHostssInfo.records;
                        vc.component.appServiceControlHostsInfo.appServiceControlHostss = _appServiceControlHostssInfo.data;
                        vc.emit('appServiceControlHosts', 'paginationPlus', 'init', {
                            total: vc.component.appServiceControlHostsInfo.records,
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
