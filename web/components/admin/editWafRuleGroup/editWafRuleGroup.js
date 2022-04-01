(function (vc, vm) {

    vc.extends({
        data: {
            editWafRuleGroupInfo: {
                groupId: '',
                groupName: '',
                state: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editWafRuleGroup', 'openEditWafRuleGroupModal', function (_params) {
                vc.component.refreshEditWafRuleGroupInfo();
                $('#editWafRuleGroupModel').modal('show');
                vc.copyObject(_params, vc.component.editWafRuleGroupInfo);
            });
        },
        methods: {
            editWafRuleGroupValidate: function () {
                return vc.validate.validate({
                    editWafRuleGroupInfo: vc.component.editWafRuleGroupInfo
                }, {
                    'editWafRuleGroupInfo.groupName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "组名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "组名称不能超过64"
                        },
                    ],
                    'editWafRuleGroupInfo.groupId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "编号不能为空"
                        }]

                });
            },
            editWafRuleGroup: function () {
                if (!vc.component.editWafRuleGroupValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/firewall/updateWafRuleGroup',
                    JSON.stringify(vc.component.editWafRuleGroupInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editWafRuleGroupModel').modal('hide');
                            vc.emit('wafRuleGroupManage', 'listWafRuleGroup', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditWafRuleGroupInfo: function () {
                vc.component.editWafRuleGroupInfo = {
                    groupId: '',
                    groupName: '',
                    state: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
