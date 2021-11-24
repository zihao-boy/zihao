/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appServiceManageInfo: {
                appServices: [],
                total: 0,
                records: 1,
                moreCondition: false,
                asId: '',
                conditions: {
                    asName: '',
                    asType: '',
                    asCount: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listAppServices(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('appServiceManage', 'listAppService', function (_param) {
                vc.component._listAppServices(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listAppServices(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listAppServices: function (_page, _rows) {

                vc.component.appServiceManageInfo.conditions.page = _page;
                vc.component.appServiceManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.appServiceManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function (json, res) {
                        var _appServiceManageInfo = JSON.parse(json);
                        vc.component.appServiceManageInfo.total = _appServiceManageInfo.total;
                        vc.component.appServiceManageInfo.records = _appServiceManageInfo.records;
                        vc.component.appServiceManageInfo.appServices = _appServiceManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.appServiceManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddAppServiceModal: function () {
                vc.emit('addAppService', 'openAddAppServiceModal', {});
            },
            _openEditAppServiceModel: function (_appService) {
                vc.emit('editAppService', 'openEditAppServiceModal', _appService);
            },
            _openDeleteAppServiceModel: function (_appService) {
                vc.emit('deleteAppService', 'openDeleteAppServiceModal', _appService);
            },
            _queryAppServiceMethod: function () {
                vc.component._listAppServices(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.appServiceManageInfo.moreCondition) {
                    vc.component.appServiceManageInfo.moreCondition = false;
                } else {
                    vc.component.appServiceManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
