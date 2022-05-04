/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            firewallRuleGroupManageInfo: {
                firewallRuleGroups: [],
                total: 0,
                records: 1,
                moreCondition: false,
                groupId: '',
                conditions: {
                    groupId: '',
                    groupName: '',

                }
            }
        },
        _initMethod: function () {
            $that._listFirewallRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('firewallRuleGroupManage', 'listFirewallRuleGroup', function (_param) {
                $that._listFirewallRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                $that._listFirewallRuleGroups(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listFirewallRuleGroups: function (_page, _rows) {

                $that.firewallRuleGroupManageInfo.conditions.page = _page;
                $that.firewallRuleGroupManageInfo.conditions.row = _rows;
                var param = {
                    params: $that.firewallRuleGroupManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/firewall/getFirewallRuleGroup',
                    param,
                    function (json, res) {
                        let _firewallRuleGroupManageInfo = JSON.parse(json);
                        $that.firewallRuleGroupManageInfo.total = _firewallRuleGroupManageInfo.total;
                        $that.firewallRuleGroupManageInfo.records = _firewallRuleGroupManageInfo.records;
                        $that.firewallRuleGroupManageInfo.firewallRuleGroups = _firewallRuleGroupManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: $that.firewallRuleGroupManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddFirewallRuleGroupModal: function () {
                vc.emit('addFirewallRuleGroup', 'openAddFirewallRuleGroupModal', {});
            },
            _openEditFirewallRuleGroupModel: function (_firewallRuleGroup) {
                vc.emit('editFirewallRuleGroup', 'openEditFirewallRuleGroupModal', _firewallRuleGroup);
            },
            _openDeleteFirewallRuleGroupModel: function (_firewallRuleGroup) {
                vc.emit('deleteFirewallRuleGroup', 'openDeleteFirewallRuleGroupModal', _firewallRuleGroup);
            },
            _queryFirewallRuleGroupMethod: function () {
                $that._listFirewallRuleGroups(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if ($that.firewallRuleGroupManageInfo.moreCondition) {
                    $that.firewallRuleGroupManageInfo.moreCondition = false;
                } else {
                    $that.firewallRuleGroupManageInfo.moreCondition = true;
                }
            },
            _openFireRulePage:function(_firewallRuleGroup){
                vc.jumpToPage('/#/pages/admin/firewallRuleManage?groupId='+_firewallRuleGroup.groupId+"&groupName="+_firewallRuleGroup.groupName)
            }


        }
    });
})(window.vc);
