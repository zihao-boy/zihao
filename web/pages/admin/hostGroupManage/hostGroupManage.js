/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostGroupManageInfo: {
                hostGroups: [],
                total: 0,
                records: 1,
                moreCondition: false,
                groupId: '',
                conditions: {
                    name: '',
                    groupId: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listHostGroups(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('hostGroupManage', 'listHostGroup', function (_param) {
                vc.component._listHostGroups(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listHostGroups(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listHostGroups: function (_page, _rows) {

                vc.component.hostGroupManageInfo.conditions.page = _page;
                vc.component.hostGroupManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.hostGroupManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/host/getHostGroup',
                    param,
                    function (json, res) {
                        var _hostGroupManageInfo = JSON.parse(json);
                        vc.component.hostGroupManageInfo.total = _hostGroupManageInfo.total;
                        vc.component.hostGroupManageInfo.records = _hostGroupManageInfo.records;
                        vc.component.hostGroupManageInfo.hostGroups = _hostGroupManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.hostGroupManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddHostGroupModal: function () {
                vc.emit('addHostGroup', 'openAddHostGroupModal', {});
            },
            _openEditHostGroupModel: function (_hostGroup) {
                vc.emit('editHostGroup', 'openEditHostGroupModal', _hostGroup);
            },
            _openDeleteHostGroupModel: function (_hostGroup) {
                vc.emit('deleteHostGroup', 'openDeleteHostGroupModal', _hostGroup);
            },
            _queryHostGroupMethod: function () {
                vc.component._listHostGroups(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.hostGroupManageInfo.moreCondition) {
                    vc.component.hostGroupManageInfo.moreCondition = false;
                } else {
                    vc.component.hostGroupManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
