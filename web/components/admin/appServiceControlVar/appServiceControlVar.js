/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            appServiceControlVarInfo: {
                appServiceControlVars: [],
                total: 0,
                records: 1,
                asId: '',
            }
        },
        _initMethod: function() {},
        _initEvent: function() {

            vc.on('appServiceControlVar', 'switch', function(_param) {
                if (_param.asId == '') {
                    return;
                }
                vc.copyObject(_param, $that.appServiceControlVarInfo)
                vc.component._listappServiceControlVars(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('appServiceControlVar', 'paginationPlus', 'page_event', function(_currentPage) {
                vc.component._listappServiceControlVars(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listappServiceControlVars: function(_page, _rows) {

                var param = {
                    params: {
                        page: _page,
                        row: _rows,
                        asId: $that.appServiceControlVarInfo.asId
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppServiceVar',
                    param,
                    function(json, res) {
                        var _appServiceControlVarsInfo = JSON.parse(json);
                        vc.component.appServiceControlVarInfo.total = _appServiceControlVarsInfo.total;
                        vc.component.appServiceControlVarInfo.records = _appServiceControlVarsInfo.records;
                        vc.component.appServiceControlVarInfo.appServiceControlVars = _appServiceControlVarsInfo.data;
                        vc.emit('appServiceControlVar', 'paginationPlus', 'init', {
                            total: vc.component.appServiceControlVarInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddVarModal: function() {
                vc.emit('addAppServiceVar', 'openAddAppServiceVarModal', $that.appServiceControlVarInfo);
            },
            _openUpdateVarModal: function(_dir) {
                vc.emit('editAppServiceVar', 'openEditAppServiceVarModal', _dir);
            },
            _openDeleteVarModal: function(_dir) {
                vc.emit('deleteAppServiceVar', 'openDeleteAppServiceVarModal', _dir);
            }

        }
    });
})(window.vc);