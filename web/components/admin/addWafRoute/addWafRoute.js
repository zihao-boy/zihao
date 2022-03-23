(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addWafRouteInfo: {
                routeId: '',
                wafId: '',
                hostname: '',
                ip: '',
                port: '80',
                wafs: []
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addWafRoute', 'openAddWafRouteModal', function () {
                $that._loadWaf();
                $('#addWafRouteModel').modal('show');
            });
        },
        methods: {
            addWafRouteValidate() {
                return vc.validate.validate({
                    addWafRouteInfo: vc.component.addWafRouteInfo
                }, {
                    'addWafRouteInfo.wafId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "Waf编号不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "Waf编号不能超过64"
                        },
                    ],
                    'addWafRouteInfo.hostname': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "域名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "域名不能超过128"
                        },
                    ],
                    'addWafRouteInfo.ip': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "应用IP不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "应用IP不能超过128"
                        },
                    ],
                    'addWafRouteInfo.port': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "应用端口不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "应用端口不能超过64"
                        },
                    ],




                });
            },
            saveWafRouteInfo: function () {
                if (!vc.component.addWafRouteValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addWafRouteInfo);
                    $('#addWafRouteModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/firewall/saveWafRoute',
                    JSON.stringify(vc.component.addWafRouteInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addWafRouteModel').modal('hide');
                            vc.component.clearAddWafRouteInfo();
                            vc.emit('wafRouteManage', 'listWafRoute', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddWafRouteInfo: function () {
                vc.component.addWafRouteInfo = {
                    wafId: '',
                    hostname: '',
                    ip: '',
                    port: '80',
                    wafs: []

                };
            },
            _loadWaf: function () {

                var param = {
                    params: {
                        page: 1,
                        row: 100
                    }
                };

                //发送get请求
                vc.http.apiGet('/firewall/getWaf',
                    param,
                    function (json, res) {
                        var _wafManageInfo = JSON.parse(json);
                        vc.component.addWafRouteInfo.wafs = _wafManageInfo.data;

                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );

            }
        }
    });

})(window.vc);
