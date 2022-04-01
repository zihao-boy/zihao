(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addWafIpBlackWhiteInfo: {
                id: '',
                typeCd: '',
                ip: '',
                scope:'*',
                seq:'',
                state:'start',
                groupId:'',
                wafRuleGroups:[]

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addWafIpBlackWhite', 'openAddWafIpBlackWhiteModal', function() {
                $that._listWafRuleGroups();
                $('#addWafIpBlackWhiteModel').modal('show');
            });
        },
        methods: {
            addWafIpBlackWhiteValidate() {
                return vc.validate.validate({
                    addWafIpBlackWhiteInfo: vc.component.addWafIpBlackWhiteInfo
                }, {
                    'addWafIpBlackWhiteInfo.typeCd': [{
                            limit: "required",
                            param: "",
                            errInfo: "类型不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "类型不能超过64"
                        },
                    ],
                    'addWafIpBlackWhiteInfo.ip': [{
                            limit: "required",
                            param: "",
                            errInfo: "IP'不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "IP'不能超过64"
                        },
                    ],




                });
            },
            saveWafIpBlackWhiteInfo: function() {
                if (!vc.component.addWafIpBlackWhiteValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addWafIpBlackWhiteInfo);
                    $('#addWafIpBlackWhiteModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/firewall/saveWafIpBlackWhite',
                    JSON.stringify(vc.component.addWafIpBlackWhiteInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addWafIpBlackWhiteModel').modal('hide');
                            vc.component.clearAddWafIpBlackWhiteInfo();
                            vc.emit('wafIpBlackWhiteManage', 'listWafIpBlackWhite', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddWafIpBlackWhiteInfo: function() {
                vc.component.addWafIpBlackWhiteInfo = {
                    typeCd: '',
                    ip: '',
                    scope:'*',
                    seq:'',
                    state:'start',
                    groupId:'',
                    wafRuleGroups:[]
                };
            },
            _listWafRuleGroups: function () {

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
                       
                        vc.component.addWafIpBlackWhiteInfo.wafRuleGroups = _wafRuleGroupManageInfo.data;
                       
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);