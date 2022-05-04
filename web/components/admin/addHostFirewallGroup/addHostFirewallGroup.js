(function(vc) {

    vc.extends({
        data: {
            addHostFirewallGroupInfo: {
                hostId: '',
                groups: [],
                groupId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addHostFirewallGroup', 'openAddHostFirewallGroupModal', function(_param) {
                $that.addHostFirewallGroupInfo.hostId = _param.hostId;
                $that._listFirewallRuleGroups();
                $('#addHostFirewallGroupModel').modal('show');
            });
        },
        methods: {
            addHostFirewallGroupValidate() {
                return vc.validate.validate({
                    addHostFirewallGroupInfo: vc.component.addHostFirewallGroupInfo
                }, {
                    'addHostFirewallGroupInfo.hostId': [{
                            limit: "required",
                            param: "",
                            errInfo: "主机不能为空"
                        },
                    ],
                    'addHostFirewallGroupInfo.groupId': [{
                            limit: "required",
                            param: "",
                            errInfo: "组不能为空"
                        },
                    ]
                });
            },
            saveHostFirewallGroupInfo: function() {
                if (!vc.component.addHostFirewallGroupValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/firewall/saveHostFirewallGroup',
                    JSON.stringify(vc.component.addHostFirewallGroupInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addHostFirewallGroupModel').modal('hide');
                            vc.emit('hostFirewallGroup', 'switch', $that.addHostFirewallGroupInfo);
                            vc.component.clearaddHostFirewallGroupInfo();
                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            clearaddHostFirewallGroupInfo: function() {
                vc.component.addHostFirewallGroupInfo = {
                    hostId: '',
                    groups: [],
                    groupId: ''
                };
            },
            _listFirewallRuleGroups: function (_page, _rows) {

                let param = {
                    params: {
                        page:1,
                        row:100
                    }
                };

                //发送get请求
                vc.http.apiGet('/firewall/getFirewallRuleGroup',
                    param,
                    function (json, res) {
                        let _firewallRuleGroupManageInfo = JSON.parse(json);
                        $that.addHostFirewallGroupInfo.groups = _firewallRuleGroupManageInfo.data;
                     
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openFireRulePage:function(_firewallRuleGroup){
                vc.jumpToPage('/#/pages/admin/firewallRuleManage?groupId='+_firewallRuleGroup.groupId+"&groupName="+_firewallRuleGroup.groupName)
            }
        }
    });

})(window.vc);