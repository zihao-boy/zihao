(function (vc, vm) {

    vc.extends({
        data: {
            editTenantInfo: {
                tenantId: '',
                tenantName: '',
                tenantType: '',
                address: '',
                personName: '',
                phone: '',
                remark: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editTenant', 'openEditTenantModal', function (_params) {
                vc.component.refreshEditTenantInfo();
                $('#editTenantModel').modal('show');
                vc.copyObject(_params, vc.component.editTenantInfo);
            });
        },
        methods: {
            editTenantValidate: function () {
                return vc.validate.validate({
                    editTenantInfo: vc.component.editTenantInfo
                }, {
                    'editTenantInfo.tenantName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "租户名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "租户名称太长"
                        },
                    ],
                    'editTenantInfo.tenantType': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "租户类型不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "12",
                            errInfo: "租户类型太长"
                        },
                    ],
                    'editTenantInfo.address': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "租户联系地址不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "256",
                            errInfo: "租户联系地址太长"
                        },
                    ],
                    'editTenantInfo.personName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "联系人不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "联系人太长"
                        },
                    ],
                    'editTenantInfo.phone': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "联系电话不能为空"
                        },
                        {
                            limit: "phone",
                            param: "11",
                            errInfo: "联系电话格式错误"
                        },
                    ],
                    'editTenantInfo.remark': [
                        {
                            limit: "maxLength",
                            param: "1000",
                            errInfo: "备注太长"
                        },
                    ],
                    'editTenantInfo.tenantId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "租户ID不能为空"
                        }]

                });
            },
            editTenant: function () {
                if (!vc.component.editTenantValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/tenant/updateTenant',
                    JSON.stringify(vc.component.editTenantInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editTenantModel').modal('hide');
                            vc.emit('tenantManage', 'listTenant', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditTenantInfo: function () {
                vc.component.editTenantInfo = {
                    tenantId: '',
                    tenantName: '',
                    tenantType: '',
                    address: '',
                    personName: '',
                    phone: '',
                    remark: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
