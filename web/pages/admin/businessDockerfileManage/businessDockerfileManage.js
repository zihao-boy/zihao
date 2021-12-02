/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            businessDockerfileManageInfo: {
                businessDockerfiles: [],
                total: 0,
                records: 1,
                moreCondition: false,
                id: '',
                conditions: {
                    id: '',
                    name: ''
                }
            }
        },
        _initMethod: function() {
            vc.component._listBusinessDockerfiles(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {
            vc.on('businessDockerfileManage', 'listBusinessDockerfile', function(_param) {
                vc.component._listBusinessDockerfiles(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listBusinessDockerfiles(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listBusinessDockerfiles: function(_page, _rows) {
                vc.component.businessDockerfileManageInfo.conditions.page = _page;
                vc.component.businessDockerfileManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.businessDockerfileManageInfo.conditions
                };
                //发送get请求
                vc.http.apiGet('/soft/getBusinessDockerfile',
                    param,
                    function(json, res) {
                        var _businessDockerfileManageInfo = JSON.parse(json);
                        vc.component.businessDockerfileManageInfo.total = _businessDockerfileManageInfo.total;
                        vc.component.businessDockerfileManageInfo.records = _businessDockerfileManageInfo.records;
                        vc.component.businessDockerfileManageInfo.businessDockerfiles = _businessDockerfileManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.businessDockerfileManageInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddBusinessDockerfileModal: function() {
                vc.emit('addBusinessDockerfile', 'openAddBusinessDockerfileModal', {});
            },
            _openEditBusinessDockerfileModel: function(_businessDockerfile) {
                vc.emit('editBusinessDockerfile', 'openEditBusinessDockerfileModal', _businessDockerfile);
            },
            _openDeleteBusinessDockerfileModel: function(_businessDockerfile) {
                vc.emit('deleteBusinessDockerfile', 'openDeleteBusinessDockerfileModal', _businessDockerfile);
            },
            _queryBusinessDockerfileMethod: function() {
                vc.component._listBusinessDockerfiles(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function() {
                if (vc.component.businessDockerfileManageInfo.moreCondition) {
                    vc.component.businessDockerfileManageInfo.moreCondition = false;
                } else {
                    vc.component.businessDockerfileManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);