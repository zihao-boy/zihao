(function (vc, vm) {

    vc.extends({
        data: {
            editMappingInfo: {
                id: '',
                domain: '',
                name: '',
                zkeys: '',
                value: '',
                remark: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editMapping', 'openEditMappingModal', function (_params) {
                vc.component.refreshEditMappingInfo();
                $('#editMappingModel').modal('show');
                vc.copyObject(_params, vc.component.editMappingInfo);
            });
        },
        methods: {
            editMappingValidate: function () {
                return vc.validate.validate({
                    editMappingInfo: vc.component.editMappingInfo
                }, {
                    'editMappingInfo.domain': [
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
                    'editMappingInfo.name': [
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
                    'editMappingInfo.zkeys': [
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
                    'editMappingInfo.value': [
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
                    'editMappingInfo.remark': [
                        {
                            limit: "maxLength",
                            param: "1000",
                            errInfo: "备注太长"
                        },
                    ],
                    'editMappingInfo.id': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "ID不能为空"
                        }]

                });
            },
            editMapping: function () {
                if (!vc.component.editMappingValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/system/updateMapping',
                    JSON.stringify(vc.component.editMappingInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editMappingModel').modal('hide');
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
            refreshEditMappingInfo: function () {
                vc.component.editMappingInfo = {
                    id: '',
                    domain: '',
                    name: '',
                    zkeys: '',
                    value: '',
                    remark: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
