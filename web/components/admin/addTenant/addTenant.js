(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addTenantInfo: {
                tenantId: '',
                tenantName: '',
                tenantType: '',
                address: '',
                personName: '',
                phone: '',
                username: '',
                passwd: '',
                remark: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addTenant', 'openAddTenantModal', function () {
                $('#addTenantModel').modal('show');
            });
        },
        methods: {
            addTenantValidate() {
                return vc.validate.validate({
                    addTenantInfo: vc.component.addTenantInfo
                }, {
                    'addTenantInfo.tenantName': [
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
                    'addTenantInfo.tenantType': [
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
                    'addTenantInfo.address': [
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
                    'addTenantInfo.personName': [
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
                    'addTenantInfo.phone': [
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
                    'addTenantInfo.username': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "登录账号不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "登录账号太长"
                        },
                    ],
                    'addTenantInfo.passwd': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "登录密码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "登录密码太长"
                        },
                    ],
                    'addTenantInfo.remark': [
                        {
                            limit: "maxLength",
                            param: "1000",
                            errInfo: "备注太长"
                        },
                    ],




                });
            },
            saveTenantInfo: function () {
                if (!vc.component.addTenantValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addTenantInfo);
                    $('#addTenantModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/tenant/saveTenant',
                    JSON.stringify(vc.component.addTenantInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addTenantModel').modal('hide');
                            vc.component.clearAddTenantInfo();
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
            clearAddTenantInfo: function () {
                vc.component.addTenantInfo = {
                    tenantName: '',
                    tenantType: '',
                    address: '',
                    personName: '',
                    phone: '',
                    username: '',
                    passwd: '',
                    remark: '',

                };
            }
        }
    });

})(window.vc);
