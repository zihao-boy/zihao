(function(vc, vm) {

    vc.extends({
        data: {
            editWafAccurateInfo: {
                id: '',
                typeCd: '',
                ip: '',
                scope: '*',
                seq: '',
                state: 'start',
                groupId: '',
                wafRuleGroups: [],
                action: '',
                typeCd: '',
                include: '',
                includeValue: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editWafAccurate', 'openEditWafAccurateModal', function(_params) {
                vc.component.refreshEditWafAccurateInfo();
                $that._listWafEditRuleGroups();
                $('#editWafAccurateModel').modal('show');
                vc.copyObject(_params, vc.component.editWafAccurateInfo);

            });
        },
        methods: {
            editWafAccurateValidate: function() {
                return vc.validate.validate({
                    editWafAccurateInfo: vc.component.editWafAccurateInfo
                }, {
                    'editWafAccurateInfo.includeValue': [{
                        limit: "required",
                        param: "",
                        errInfo: "内容不能为空"
                    }],
                    'editWafAccurateInfo.id': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editWafAccurate: function() {
                if (!vc.component.editWafAccurateValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }


                vc.http.apiPost(
                    '/firewall/updateWafAccurate',
                    JSON.stringify(vc.component.editWafAccurateInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editWafAccurateModel').modal('hide');
                            vc.emit('wafAccurateManage', 'listWafAccurate', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditWafAccurateInfo: function() {
                vc.component.editWafAccurateInfo = {
                    id: '',
                    typeCd: '',
                    ip: '',
                    scope: '*',
                    seq: '',
                    state: 'start',
                    groupId: '',
                    wafRuleGroups: [],
                    action: '',
                    typeCd: '',
                    include: '',
                    includeValue: ''

                }
            },
            _listWafEditRuleGroups: function() {

                var param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafRuleGroup',
                    param,
                    function(json, res) {
                        var _wafRuleGroupManageInfo = JSON.parse(json);

                        vc.component.editWafAccurateInfo.wafRuleGroups = _wafRuleGroupManageInfo.data;

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc, window.vc.component);