(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addDnsMapInfo: {
                dnsMapId: '',
                host: '',
                type: '',
                value: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addDnsMap', 'openAddDnsMapModal', function () {
                $('#addDnsMapModel').modal('show');
            });
        },
        methods: {
            addDnsMapValidate() {
                return vc.validate.validate({
                    addDnsMapInfo: vc.component.addDnsMapInfo
                }, {
                    'addDnsMapInfo.host': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "域名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "域名不能超过64"
                        },
                    ],
                    'addDnsMapInfo.type': [
                        {
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
                    'addDnsMapInfo.value': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "ip'不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "ip'不能超过64"
                        },
                    ],




                });
            },
            saveDnsMapInfo: function () {
                if (!vc.component.addDnsMapValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addDnsMapInfo);
                    $('#addDnsMapModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/dns/saveDnsMap',
                    JSON.stringify(vc.component.addDnsMapInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addDnsMapModel').modal('hide');
                            vc.component.clearAddDnsMapInfo();
                            vc.emit('dnsMapManage', 'listDnsMap', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddDnsMapInfo: function () {
                vc.component.addDnsMapInfo = {
                    host: '',
                    type: '',
                    value: '',

                };
            }
        }
    });

})(window.vc);
