(function (vc, vm) {

    vc.extends({
        data: {
            editFirewallRuleInfo: {
                ruleId: '',
                groupId: '',
                inout: '',
                allowLimit: '',
                seq: '',
                protocol: '',
                srcObj: '',
                dstObj: '',
                remark: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editFirewallRule', 'openEditFirewallRuleModal', function (_params) {
                vc.component.refreshEditFirewallRuleInfo();
                $('#editFirewallRuleModel').modal('show');
                vc.copyObject(_params, vc.component.editFirewallRuleInfo);
            });
        },
        methods: {
            editFirewallRuleValidate: function () {
                return vc.validate.validate({
                    editFirewallRuleInfo: vc.component.editFirewallRuleInfo
                }, {
                    'editFirewallRuleInfo.groupId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "组编号不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "组编号不能超过64"
                        },
                    ],
                    'editFirewallRuleInfo.inout': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "入站出站不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "入站出站不能超过64"
                        },
                    ],
                    'editFirewallRuleInfo.allowLimit': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "授权策略不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "授权策略不能超过64"
                        },
                    ],
                    'editFirewallRuleInfo.seq': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "顺序不能为空"
                        },
                        {
                            limit: "num",
                            param: "11",
                            errInfo: "顺序不能超过11"
                        },
                    ],
                    'editFirewallRuleInfo.protocol': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "协议不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "协议不能超过64"
                        },
                    ],
                    'editFirewallRuleInfo.srcObj': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "授权对象不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "授权对象不能超过64"
                        },
                    ],
                    'editFirewallRuleInfo.dstObj': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "端口不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "端口不能超过64"
                        },
                    ],
                    'editFirewallRuleInfo.remark': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "备注不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "512",
                            errInfo: "备注不能超过512"
                        },
                    ],
                    'editFirewallRuleInfo.ruleId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "编号不能为空"
                        }]

                });
            },
            editFirewallRule: function () {
                if (!vc.component.editFirewallRuleValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/firewall/updateFirewallRule',
                    JSON.stringify(vc.component.editFirewallRuleInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editFirewallRuleModel').modal('hide');
                            vc.emit('firewallRuleManage', 'listFirewallRule', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditFirewallRuleInfo: function () {
                vc.component.editFirewallRuleInfo = {
                    ruleId: '',
                    groupId: '',
                    inout: '',
                    allowLimit: '',
                    seq: '',
                    protocol: '',
                    srcObj: '',
                    dstObj: '',
                    remark: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
