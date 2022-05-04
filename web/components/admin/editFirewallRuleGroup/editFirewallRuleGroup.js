(function (vc, vm) {

    vc.extends({
        data: {
            editFirewallRuleGroupInfo: {
                groupId: '',
                groupName: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editFirewallRuleGroup', 'openEditFirewallRuleGroupModal', function (_params) {
                vc.component.refreshEditFirewallRuleGroupInfo();
                $('#editFirewallRuleGroupModel').modal('show');
                vc.copyObject(_params, vc.component.editFirewallRuleGroupInfo);
            });
        },
        methods: {
            editFirewallRuleGroupValidate: function () {
                return vc.validate.validate({
                    editFirewallRuleGroupInfo: vc.component.editFirewallRuleGroupInfo
                }, {
                    'editFirewallRuleGroupInfo.groupName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "名称不能超过64"
                        },
                    ],
                    'editFirewallRuleGroupInfo.groupId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "编号不能为空"
                        }]

                });
            },
            editFirewallRuleGroup: function () {
                if (!vc.component.editFirewallRuleGroupValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/firewall/updateFirewallRuleGroup',
                    JSON.stringify(vc.component.editFirewallRuleGroupInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editFirewallRuleGroupModel').modal('hide');
                            vc.emit('firewallRuleGroupManage', 'listFirewallRuleGroup', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditFirewallRuleGroupInfo: function () {
                vc.component.editFirewallRuleGroupInfo = {
                    groupId: '',
                    groupName: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
