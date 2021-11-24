(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addMappingInfo: {
                id: '',
                domain: 'DOMAIN.COMMON',
                name: '',
                zkeys: '',
                value: '',
                remark: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addMapping', 'openAddMappingModal', function () {
                $('#addMappingModel').modal('show');
            });
        },
        methods: {
            addMappingValidate() {
                return vc.validate.validate({
                    addMappingInfo: vc.component.addMappingInfo
                }, {
                    'addMappingInfo.domain': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "域不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "50",
                            errInfo: "域太长"
                        },
                    ],
                    'addMappingInfo.name': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "50",
                            errInfo: "名称太长"
                        },
                    ],
                    'addMappingInfo.zkeys': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "键不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "100",
                            errInfo: "键太长"
                        },
                    ],
                    'addMappingInfo.value': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "值不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "1000",
                            errInfo: "值太长"
                        },
                    ],
                    'addMappingInfo.remark': [
                        {
                            limit: "maxLength",
                            param: "1000",
                            errInfo: "备注太长"
                        },
                    ],




                });
            },
            saveMappingInfo: function () {
                if (!vc.component.addMappingValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }
                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addMappingInfo);
                    $('#addMappingModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/system/saveMapping',
                    JSON.stringify(vc.component.addMappingInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addMappingModel').modal('hide');
                            vc.component.clearAddMappingInfo();
                            vc.emit('mappingManage', 'listMapping', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddMappingInfo: function () {
                vc.component.addMappingInfo = {
                    domain: 'DOMAIN.COMMON',
                    name: '',
                    zkeys: '',
                    value: '',
                    remark: '',

                };
            }
        }
    });

})(window.vc);
