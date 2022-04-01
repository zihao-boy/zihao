(function(vc, vm) {

    vc.extends({
        data: {
            editWafIpBlackWhiteInfo: {
                id: '',
                typeCd: '',
                ip: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editWafIpBlackWhite', 'openEditWafIpBlackWhiteModal', function(_params) {
                vc.component.refreshEditWafIpBlackWhiteInfo();
                $('#editWafIpBlackWhiteModel').modal('show');
                vc.copyObject(_params, vc.component.editWafIpBlackWhiteInfo);
            });
        },
        methods: {
            editWafIpBlackWhiteValidate: function() {
                return vc.validate.validate({
                    editWafIpBlackWhiteInfo: vc.component.editWafIpBlackWhiteInfo
                }, {
                    'editWafIpBlackWhiteInfo.typeCd': [{
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
                    'editWafIpBlackWhiteInfo.ip': [{
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
                    'editWafIpBlackWhiteInfo.id': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editWafIpBlackWhite: function() {
                if (!vc.component.editWafIpBlackWhiteValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/firewall/updateWafIpBlackWhite',
                    JSON.stringify(vc.component.editWafIpBlackWhiteInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editWafIpBlackWhiteModel').modal('hide');
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
            refreshEditWafIpBlackWhiteInfo: function() {
                vc.component.editWafIpBlackWhiteInfo = {
                    id: '',
                    typeCd: '',
                    ip: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);