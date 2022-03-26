(function(vc, vm) {

    vc.extends({
        data: {
            editWafRouteInfo: {
                routeId: '',
                wafId: '',
                hostname: '',
                scheme: '',
                ip: '',
                port: '',
                privKeyContent: '',
                certContent: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editWafRoute', 'openEditWafRouteModal', function(_params) {
                vc.component.refreshEditWafRouteInfo();
                $('#editWafRouteModel').modal('show');
                vc.copyObject(_params, vc.component.editWafRouteInfo);
            });
        },
        methods: {
            editWafRouteValidate: function() {
                return vc.validate.validate({
                    editWafRouteInfo: vc.component.editWafRouteInfo
                }, {
                    'editWafRouteInfo.wafId': [{
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
                    'editWafRouteInfo.hostname': [{
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
                    'editWafRouteInfo.ip': [{
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
                    'editWafRouteInfo.port': [{
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
                    'editWafRouteInfo.routeId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editWafRoute: function() {
                if (!vc.component.editWafRouteValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/firewall/updateWafRoute',
                    JSON.stringify(vc.component.editWafRouteInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editWafRouteModel').modal('hide');
                            vc.emit('wafRouteManage', 'listWafRoute', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditWafRouteInfo: function() {
                vc.component.editWafRouteInfo = {
                    routeId: '',
                    wafId: '',
                    scheme: '',
                    hostname: '',
                    privKeyContent: '',
                    certContent: '',
                    ip: '',
                    port: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);