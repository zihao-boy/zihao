(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addTenantSettingInfo: {
                settingId: '',
                specCds:[],
                specCd: '',
                value: '',

            }
        },
        _initMethod: function () {

            vc.getDict('tenant_setting','spec_cd',function(data){
                
                $that.addTenantSettingInfo.specCds = data.data;
            })

        },
        _initEvent: function () {
            vc.on('addTenantSetting', 'openAddTenantSettingModal', function () {
                $('#addTenantSettingModel').modal('show');
            });
        },
        methods: {
            addTenantSettingValidate() {
                return vc.validate.validate({
                    addTenantSettingInfo: vc.component.addTenantSettingInfo
                }, {
                    'addTenantSettingInfo.specCd': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "规格不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "规格错误"
                        },
                    ],
                    'addTenantSettingInfo.value': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "值不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "1024",
                            errInfo: "值太长"
                        },
                    ],
                });
            },
            saveTenantSettingInfo: function () {
                if (!vc.component.addTenantSettingValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }
                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addTenantSettingInfo);
                    $('#addTenantSettingModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/tenant/saveTenantSetting',
                    JSON.stringify(vc.component.addTenantSettingInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addTenantSettingModel').modal('hide');
                            vc.component.clearAddTenantSettingInfo();
                            vc.emit('tenantSettingManage', 'listTenantSetting', {});

                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            clearAddTenantSettingInfo: function () {
                let _specCds = $that.addTenantSettingInfo.specCds
                vc.component.addTenantSettingInfo = {
                    specCd: '',
                    value: '',
                    specCds: _specCds
                };
            }
        }
    });

})(window.vc);
