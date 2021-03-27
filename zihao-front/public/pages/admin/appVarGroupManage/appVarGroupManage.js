/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appVarGroupManageInfo: {
                appVarGroups: [],
                total: 0,
                records: 1,
                moreCondition: false,
                avgId: '',
                conditions: {
                    avgName: '',
                    avgType: '',
                    avgId: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listAppVarGroups(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('appVarGroupManage', 'listAppVarGroup', function (_param) {
                vc.component._listAppVarGroups(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listAppVarGroups(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listAppVarGroups: function (_page, _rows) {

                vc.component.appVarGroupManageInfo.conditions.page = _page;
                vc.component.appVarGroupManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.appVarGroupManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppVarGroup',
                    param,
                    function (json, res) {
                        var _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.appVarGroupManageInfo.total = _appVarGroupManageInfo.total;
                        vc.component.appVarGroupManageInfo.records = _appVarGroupManageInfo.records;
                        vc.component.appVarGroupManageInfo.appVarGroups = _appVarGroupManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.appVarGroupManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddAppVarGroupModal: function () {
                vc.emit('addAppVarGroup', 'openAddAppVarGroupModal', {});
            },
            _openEditAppVarGroupModel: function (_appVarGroup) {
                vc.emit('editAppVarGroup', 'openEditAppVarGroupModal', _appVarGroup);
            },
            _openDeleteAppVarGroupModel: function (_appVarGroup) {
                vc.emit('deleteAppVarGroup', 'openDeleteAppVarGroupModal', _appVarGroup);
            },
            _queryAppVarGroupMethod: function () {
                vc.component._listAppVarGroups(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.appVarGroupManageInfo.moreCondition) {
                    vc.component.appVarGroupManageInfo.moreCondition = false;
                } else {
                    vc.component.appVarGroupManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
