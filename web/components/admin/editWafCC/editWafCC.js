(function(vc, vm) {

    vc.extends({
        data: {
            editWafCCInfo: {
                id: '',
                typeCd: '',
                ip: '',
                scope:'*',
                seq:'',
                state:'start',
                groupId:'',
                wafRuleGroups:[],
                visitCount: '',
                visitSec: '',
                blockSec: '',
                path: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editWafCC', 'openEditWafCCModal', function(_params) {
                vc.component.refreshEditWafCCInfo();
                $that._listWafEditRuleGroups();
                $('#editWafCCModel').modal('show');
                vc.copyObject(_params, vc.component.editWafCCInfo);
                
            });
        },
        methods: {
            editWafCCValidate: function() {
                return vc.validate.validate({
                    editWafCCInfo: vc.component.editWafCCInfo
                }, {
                    'editWafCCInfo.visitCount': [{
                        limit: "required",
                        param: "",
                        errInfo: "访问次数不能为空"
                    }],
                    'editWafCCInfo.id': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editWafCC: function() {
                if (!vc.component.editWafCCValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }


                vc.http.apiPost(
                    '/firewall/updateWafCC',
                    JSON.stringify(vc.component.editWafCCInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editWafCCModel').modal('hide');
                            vc.emit('wafCCManage', 'listWafCC', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditWafCCInfo: function() {
                vc.component.editWafCCInfo = {
                    id: '',
                    typeCd: '',
                    ip: '',
                    scope:'*',
                    seq:'',
                    state:'start',
                    groupId:'',
                    wafRuleGroups:[],
                    visitCount: '',
                    visitSec: '',
                    blockSec: '',
                    path: ''

                }
            },
            _listWafEditRuleGroups: function () {

                var param = {
                    params: {
                        page:1,
                        row:100
                    }
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafRuleGroup',
                    param,
                    function (json, res) {
                        var _wafRuleGroupManageInfo = JSON.parse(json);
                       
                        vc.component.editWafCCInfo.wafRuleGroups = _wafRuleGroupManageInfo.data;
                       
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc, window.vc.component);