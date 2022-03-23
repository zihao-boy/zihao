(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addWafInfo: {
                wafId: '',
                wafName: '',
                port: '',
                state: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addWaf', 'openAddWafModal', function () {
                $('#addWafModel').modal('show');
            });
        },
        methods: {
            addWafValidate() {
                return vc.validate.validate({
                    addWafInfo: vc.component.addWafInfo
                }, {
                    'addWafInfo.wafName': [
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
                    'addWafInfo.port': [
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
                    'addWafInfo.state': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "状态不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "12",
                            errInfo: "状态不能超过12"
                        },
                    ],




                });
            },
            saveWafInfo: function () {
                if (!vc.component.addWafValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addWafInfo);
                    $('#addWafModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/firewall/saveWaf',
                    JSON.stringify(vc.component.addWafInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addWafModel').modal('hide');
                            vc.component.clearAddWafInfo();
                            vc.emit('wafManage', 'listWaf', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddWafInfo: function () {
                vc.component.addWafInfo = {
                    wafName: '',
                    port: '',
                    state: '',

                };
            }
        }
    });

})(window.vc);
