/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            innerNetUserManageInfo: {
                innerNetUsers: [],
                total: 0,
                records: 1,
                moreCondition: false,
                routeId: '',
                conditions: {
                    routeId: '',
                    hostname: '',
                    ip: '',
                    port: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listInnerNetUsers(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('innerNetUserManage', 'listInnerNetUser', function (_param) {
                vc.component._listInnerNetUsers(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listInnerNetUsers(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listInnerNetUsers: function (_page, _rows) {

                vc.component.innerNetUserManageInfo.conditions.page = _page;
                vc.component.innerNetUserManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.innerNetUserManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/innerNet/getInnerNetUser',
                    param,
                    function (json, res) {
                        var _innerNetUserManageInfo = JSON.parse(json);
                        vc.component.innerNetUserManageInfo.total = _innerNetUserManageInfo.total;
                        vc.component.innerNetUserManageInfo.records = _innerNetUserManageInfo.records;
                        vc.component.innerNetUserManageInfo.innerNetUsers = _innerNetUserManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.innerNetUserManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddInnerNetUserModal: function () {
                vc.emit('addInnerNetUser', 'openAddInnerNetUserModal', {});
            },
            _openEditInnerNetUserModel: function (_innerNetUser) {
                vc.emit('editInnerNetUser', 'openEditInnerNetUserModal', _innerNetUser);
            },
            _openDeleteInnerNetUserModel: function (_innerNetUser) {
                vc.emit('deleteInnerNetUser', 'openDeleteInnerNetUserModal', _innerNetUser);
            },
            _queryInnerNetUserMethod: function () {
                vc.component._listInnerNetUsers(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.innerNetUserManageInfo.moreCondition) {
                    vc.component.innerNetUserManageInfo.moreCondition = false;
                } else {
                    vc.component.innerNetUserManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
