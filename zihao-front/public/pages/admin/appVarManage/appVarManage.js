/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appVarManageInfo: {
                appVars: [],
                total: 0,
                records: 1,
                moreCondition: false,
                avId: '',
                appVarGroups:[],
                conditions: {
                    varName: '',
                    avgId: '',
                    avId: '',

                }
            }
        },
        _initMethod: function () {
            $that._listManageAppVarGroups();
            vc.component._listAppVars(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('appVarManage', 'listAppVar', function (_param) {
                vc.component._listAppVars(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listAppVars(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listAppVars: function (_page, _rows) {

                vc.component.appVarManageInfo.conditions.page = _page;
                vc.component.appVarManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.appVarManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppVar',
                    param,
                    function (json, res) {
                        var _appVarManageInfo = JSON.parse(json);
                        vc.component.appVarManageInfo.total = _appVarManageInfo.total;
                        vc.component.appVarManageInfo.records = _appVarManageInfo.records;
                        vc.component.appVarManageInfo.appVars = _appVarManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.appVarManageInfo.records,
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
            _openEditAppVarModel: function (_appVar) {
                vc.emit('editAppVar', 'openEditAppVarModal', _appVar);
            },
            _openDeleteAppVarModel: function (_appVar) {
                vc.emit('deleteAppVar', 'openDeleteAppVarModal', _appVar);
            },
            _queryAppVarMethod: function () {
                vc.component._listAppVars(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.appVarManageInfo.moreCondition) {
                    vc.component.appVarManageInfo.moreCondition = false;
                } else {
                    vc.component.appVarManageInfo.moreCondition = true;
                }
            },
            _listManageAppVarGroups: function () {

                var param = {
                    params: {
                        page:1,
                        row:50
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppVarGroup',
                    param,
                    function (json, res) {
                        let _appVarGroupManageInfo = JSON.parse(json);
                        vc.component.appVarManageInfo.appVarGroups = _appVarGroupManageInfo.data;   
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },


        }
    });
})(window.vc);
