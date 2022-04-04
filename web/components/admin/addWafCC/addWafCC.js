(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addWafCCInfo: {
                id: '',
                typeCd: '',
                ip: '',
                scope: '*',
                seq: '',
                state: 'start',
                groupId: '',
                wafRuleGroups: [],
                visitCount: '',
                visitSec: '',
                blockSec: '',
                path: ''


            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addWafCC', 'openAddWafCCModal', function () {
                $that._listWafRuleGroups();
                $('#addWafCCModel').modal('show');
            });
        },
        methods: {
            addWafCCValidate() {
                return vc.validate.validate({
                    addWafCCInfo: vc.component.addWafCCInfo
                }, {
                    'addWafCCInfo.visitCount': [{
                        limit: "required",
                        param: "",
                        errInfo: "访问次数不能为空"
                    }
                    ],

                });
            },
            saveWafCCInfo: function () {
                if (!vc.component.addWafCCValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addWafCCInfo);
                    $('#addWafCCModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/firewall/saveWafCC',
                    JSON.stringify(vc.component.addWafCCInfo), {
                    emulateJSON: true
                },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addWafCCModel').modal('hide');
                            vc.component.clearAddWafCCInfo();
                            vc.emit('wafCCManage', 'listWafCC', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddWafCCInfo: function () {
                vc.component.addWafCCInfo = {
                    typeCd: '',
                    ip: '',
                    scope: '*',
                    seq: '',
                    state: 'start',
                    groupId: '',
                    wafRuleGroups: [],
                    visitCount: '',
                    visitSec: '',
                    blockSec: '',
                    path: ''
                };
            },
            _listWafRuleGroups: function () {

                var param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWafRuleGroup',
                    param,
                    function (json, res) {
                        var _wafRuleGroupManageInfo = JSON.parse(json);

                        vc.component.addWafCCInfo.wafRuleGroups = _wafRuleGroupManageInfo.data;

                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);