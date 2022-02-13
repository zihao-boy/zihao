(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addOssInfo: {
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
            vc.on('addOss', 'openAddOssModal', function() {
                $('#addOssModel').modal('show');
            });
        },
        methods: {
            addOssValidate() {
                return vc.validate.validate({
                    addOssInfo: vc.component.addOssInfo
                }, {
                    'addOssInfo.name': [{
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
                    'addOssInfo.ossType': [{
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
                    'addOssInfo.bucket': [{
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
                    'addOssInfo.accessKeySecret': [{
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
                    'addOssInfo.accessKeyId': [{
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
                    'addOssInfo.endpoint': [{
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
                    'addOssInfo.path': [{
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




                });
            },
            saveOssInfo: function() {
                if (!vc.component.addOssValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addOssInfo);
                    $('#addOssModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/resources/saveOss',
                    JSON.stringify(vc.component.addOssInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addOssModel').modal('hide');
                            vc.component.clearAddOssInfo();
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
            clearAddOssInfo: function() {
                vc.component.addOssInfo = {
                    name: '',
                    ossType: '',
                    bucket: '',
                    accessKeySecret: '',
                    accessKeyId: '',
                    endpoint: '',
                    path: '',

                };
            }
        }
    });

})(window.vc);