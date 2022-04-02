/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            wafRuleGroupManageInfo: {
                wafRuleGroups: [],
                total: 0,
                records: 1,
                moreCondition: false,
                groupId: '',
                conditions: {
                    groupId: '',
                    groupName: '',
                    state: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listWafRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('wafRuleGroupManage', 'listWafRuleGroup', function (_param) {
                vc.component._listWafRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listWafRuleGroups(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listWafRuleGroups: function (_page, _rows) {

                vc.component.wafRuleGroupManageInfo.conditions.page = _page;
                vc.component.wafRuleGroupManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.wafRuleGroupManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafRuleGroup',
                    param,
                    function (json, res) {
                        var _wafRuleGroupManageInfo = JSON.parse(json);
                        vc.component.wafRuleGroupManageInfo.total = _wafRuleGroupManageInfo.total;
                        vc.component.wafRuleGroupManageInfo.records = _wafRuleGroupManageInfo.records;
                        vc.component.wafRuleGroupManageInfo.wafRuleGroups = _wafRuleGroupManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.wafRuleGroupManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddWafRuleGroupModal: function () {
                vc.emit('addWafRuleGroup', 'openAddWafRuleGroupModal', {});
            },
            _openEditWafRuleGroupModel: function (_wafRuleGroup) {
                vc.emit('editWafRuleGroup', 'openEditWafRuleGroupModal', _wafRuleGroup);
            },
            _openDeleteWafRuleGroupModel: function (_wafRuleGroup) {
                vc.emit('deleteWafRuleGroup', 'openDeleteWafRuleGroupModal', _wafRuleGroup);
            },
            _queryWafRuleGroupMethod: function () {
                vc.component._listWafRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.wafRuleGroupManageInfo.moreCondition) {
                    vc.component.wafRuleGroupManageInfo.moreCondition = false;
                } else {
                    vc.component.wafRuleGroupManageInfo.moreCondition = true;
                }
            },
            _openToRulePage:function(_wafRuleGroup){
                vc.jumpToPage('/index.html#/pages/admin/wafRuleManage?groupId='+_wafRuleGroup.groupId);
            },
            _openStartRuleGroupModel:function(_wafRuleGroup){
                vc.http.apiPost(
                    '/firewall/startWafRuleGroup',
                    JSON.stringify(_wafRuleGroup),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vc.toast('启动成功');
                            vc.emit('wafRuleGroupManage', 'listWafRuleGroup', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            }


        }
    });
})(window.vc);
