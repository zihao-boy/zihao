/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            tenantManageInfo: {
                tenants: [],
                total: 0,
                records: 1,
                moreCondition: false,
                tenantId: '',
                conditions: {
                    tenantName: '',
                    tenantType: '',
                    personName: '',
                }
            }
        },
        _initMethod: function () {
            vc.component._listTenants(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('tenantManage', 'listTenant', function (_param) {
                vc.component._listTenants(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listTenants(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listTenants: function (_page, _rows) {

                vc.component.tenantManageInfo.conditions.page = _page;
                vc.component.tenantManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.tenantManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/tenant/getTenants',
                    param,
                    function (json, res) {
                        var _tenantManageInfo = JSON.parse(json);
                        vc.component.tenantManageInfo.total = _tenantManageInfo.total;
                        vc.component.tenantManageInfo.records = _tenantManageInfo.records;
                        vc.component.tenantManageInfo.tenants = _tenantManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.tenantManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddTenantModal: function () {
                vc.emit('addTenant', 'openAddTenantModal', {});
            },
            _openEditTenantModel: function (_tenant) {
                vc.emit('editTenant', 'openEditTenantModal', _tenant);
            },
            _openDeleteTenantModel: function (_tenant) {
                vc.emit('deleteTenant', 'openDeleteTenantModal', _tenant);
            },
            _queryTenantMethod: function () {
                vc.component._listTenants(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.tenantManageInfo.moreCondition) {
                    vc.component.tenantManageInfo.moreCondition = false;
                } else {
                    vc.component.tenantManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
