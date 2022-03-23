(function (vc, vm) {

    vc.extends({
        data: {
            editWafInfo: {
                wafId: '',
                wafName: '',
                port: '',
                state: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editWaf', 'openEditWafModal', function (_params) {
                vc.component.refreshEditWafInfo();
                $('#editWafModel').modal('show');
                vc.copyObject(_params, vc.component.editWafInfo);
            });
        },
        methods: {
            editWafValidate: function () {
                return vc.validate.validate({
                    editWafInfo: vc.component.editWafInfo
                }, {
                    'editWafInfo.wafName': [
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
                    'editWafInfo.port': [
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
                    'editWafInfo.state': [
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
                    'editWafInfo.wafId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "编号不能为空"
                        }]

                });
            },
            editWaf: function () {
                if (!vc.component.editWafValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/firewall/updateWaf',
                    JSON.stringify(vc.component.editWafInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editWafModel').modal('hide');
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
            refreshEditWafInfo: function () {
                vc.component.editWafInfo = {
                    wafId: '',
                    wafName: '',
                    port: '',
                    state: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
