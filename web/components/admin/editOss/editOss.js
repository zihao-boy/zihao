(function(vc, vm) {

    vc.extends({
        data: {
            editOssInfo: {
                ossId: '',
                name: '',
                ossType: '',
                bucket: '',
                accessKeySecret: '',
                accessKeyId: '',
                endpoint: '',
                path: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editOss', 'openEditOssModal', function(_params) {
                vc.component.refreshEditOssInfo();
                $('#editOssModel').modal('show');
                vc.copyObject(_params, vc.component.editOssInfo);
            });
        },
        methods: {
            editOssValidate: function() {
                return vc.validate.validate({
                    editOssInfo: vc.component.editOssInfo
                }, {
                    'editOssInfo.name': [{
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
                    'editOssInfo.ossType': [{
                            limit: "required",
                            param: "",
                            errInfo: "类型不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "12",
                            errInfo: "类型不能超过12"
                        },
                    ],
                    'editOssInfo.bucket': [{
                            limit: "required",
                            param: "",
                            errInfo: "bucket不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "bucket不能超过128"
                        },
                    ],
                    'editOssInfo.accessKeySecret': [{
                            limit: "required",
                            param: "",
                            errInfo: "access_key_secret不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "access_key_secret不能超过128"
                        },
                    ],
                    'editOssInfo.accessKeyId': [{
                            limit: "required",
                            param: "",
                            errInfo: "access_key_id不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "access_key_id不能超过128"
                        },
                    ],
                    'editOssInfo.endpoint': [{
                            limit: "required",
                            param: "",
                            errInfo: "endpoint不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "endpoint不能超过128"
                        },
                    ],
                    'editOssInfo.path': [{
                            limit: "required",
                            param: "",
                            errInfo: "路劲不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "路劲不能超过128"
                        },
                    ],
                    'editOssInfo.ossId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editOss: function() {
                if (!vc.component.editOssValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/resources/updateOss',
                    JSON.stringify(vc.component.editOssInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editOssModel').modal('hide');
                            vc.emit('ossManage', 'listOss', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditOssInfo: function() {
                vc.component.editOssInfo = {
                    ossId: '',
                    name: '',
                    ossType: '',
                    bucket: '',
                    accessKeySecret: '',
                    accessKeyId: '',
                    endpoint: '',
                    path: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);