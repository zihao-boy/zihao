(function(vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addBusinessPackageInfo: {
                id: '',
                name: '',
                varsion: '',
                excelTemplate: '',
                type: '1001',
                filename: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addBusinessPackage', 'openAddBusinessPackageModal', function() {
                $('#addBusinessPackageModel').modal('show');
            });
        },
        methods: {
            addBusinessPackageValidate() {
                return vc.validate.validate({
                    addBusinessPackageInfo: vc.component.addBusinessPackageInfo
                }, {
                    'addBusinessPackageInfo.name': [{
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "名称不能超过64"
                        },
                    ]
                });
            },
            saveBusinessPackageInfo: function() {
                if (!vc.component.addBusinessPackageValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                var param = new FormData();
                param.append("uploadFile", vc.component.addBusinessPackageInfo.excelTemplate);
                param.append('name', vc.component.addBusinessPackageInfo.name);
                param.append('type', vc.component.addBusinessPackageInfo.type);
                param.append('filename', vc.component.addBusinessPackageInfo.filename);



                vc.http.apiPost(
                    '/soft/saveBusinessPackages',
                    param, {
                        emulateJSON: true,
                        //添加请求头
                        headers: {
                            "Content-Type": "multipart/form-data"
                        }
                    },
                    function(json, res) {
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addBusinessPackageModel').modal('hide');
                            vc.component.clearAddBusinessPackageInfo();
                            vc.emit('businessPackageManage', 'listBusinessPackage', {});

                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            getExcelTemplate: function(e) {
                //console.log("getExcelTemplate 开始调用")
                vc.component.addBusinessPackageInfo.excelTemplate = e.target.files[0];
            },
            clearAddBusinessPackageInfo: function() {
                vc.component.addBusinessPackageInfo = {
                    name: '',
                    varsion: '',
                    excelTemplate: '',
                    type: '1001',
                    filename: ''
                };
            }
        }
    });

})(window.vc);