(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addWafAccurateInfo: {
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
            vc.on('addWafAccurate', 'openAddWafAccurateModal', function() {
                $that._listWafRuleGroups();
                $('#addWafAccurateModel').modal('show');
            });
        },
        methods: {
            addWafAccurateValidate() {
                return vc.validate.validate({
                    addWafAccurateInfo: vc.component.addWafAccurateInfo
                }, {
                    'addWafAccurateInfo.includeValue': [{
                        limit: "required",
                        param: "",
                        errInfo: "内容不能为空"
                    }],

                });
            },
            saveWafAccurateInfo: function() {
                if (!vc.component.addWafAccurateValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addWafAccurateInfo);
                    $('#addWafAccurateModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/firewall/saveWafAccurate',
                    JSON.stringify(vc.component.addWafAccurateInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addWafAccurateModel').modal('hide');
                            vc.component.clearAddWafAccurateInfo();
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
            clearAddWafAccurateInfo: function() {
                vc.component.addWafAccurateInfo = {
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
                };
            },
            _listWafRuleGroups: function() {

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

                        vc.component.addWafAccurateInfo.wafRuleGroups = _wafRuleGroupManageInfo.data;

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);