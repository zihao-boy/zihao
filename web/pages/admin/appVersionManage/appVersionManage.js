/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appVersionManageInfo: {
                total: 0,
                records: 1,
                moreCondition: false,
                avId: '',
                appVersions:[],
                conditions: {
                    name: '',
                    avgId: '',
                    avId: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listAppVersions(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {
            vc.on('appVersionManage', 'listAppVar', function (_param) {
                vc.component._listAppVersions(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listAppVersions(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listAppVersions: function (_page, _rows) {

                vc.component.appVersionManageInfo.conditions.page = _page;
                vc.component.appVersionManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.appVersionManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/appVersion/getAppVersion',
                    param,
                    function (json, res) {
                        var _appVersionManageInfo = JSON.parse(json);
                        vc.component.appVersionManageInfo.total = _appVersionManageInfo.total;
                        vc.component.appVersionManageInfo.records = _appVersionManageInfo.records;
                        vc.component.appVersionManageInfo.appVersions = _appVersionManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.appVersionManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddAppVarModal: function () {
                vc.emit('addAppVar', 'openAddAppVarModal', {});
            },
            _queryAppVarMethod: function () {
                vc.component._listAppVersions(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function () {
                if (vc.component.appVersionManageInfo.moreCondition) {
                    vc.component.appVersionManageInfo.moreCondition = false;
                } else {
                    vc.component.appVersionManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
