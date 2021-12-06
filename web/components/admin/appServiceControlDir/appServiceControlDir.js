/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appServiceControlDirInfo: {
                appServiceControlDirs: [],
                total: 0,
                records: 1,
                asId: '',
            }
        },
        _initMethod: function () {
        },
        _initEvent: function () {

            vc.on('appServiceControlDir', 'switch', function (_param) {
                if (_param.asId == '') {
                    return;
                }
                vc.copyObject(_param, $that.appServiceControlDirInfo)
                vc.component._listappServiceControlDirs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('appServiceControlDir', 'paginationPlus', 'page_event', function (_currentPage) {
                vc.component._listappServiceControlDirs(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listappServiceControlDirs: function (_page, _rows) {

                var param = {
                    params: {
                        page: _page,
                        row: _rows,
                        asId: $that.appServiceControlDirInfo.asId
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppServiceDir',
                    param,
                    function (json, res) {
                        var _appServiceControlDirsInfo = JSON.parse(json);
                        vc.component.appServiceControlDirInfo.total = _appServiceControlDirsInfo.total;
                        vc.component.appServiceControlDirInfo.records = _appServiceControlDirsInfo.records;
                        vc.component.appServiceControlDirInfo.appServiceControlDirs = _appServiceControlDirsInfo.data;
                        vc.emit('appServiceControlDir', 'paginationPlus', 'init', {
                            total: vc.component.appServiceControlDirInfo.records,
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
