(function(vc, vm) {

    vc.extends({
        data: {
            editBusinessPackageInfo: {
                id: '',
                name: '',
                varsion: '',
                excelEditTemplate: '',
                type: '1001',
                filename: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editBusinessPackage', 'openEditBusinessPackageModal', function(_params) {
                vc.component.refreshEditBusinessPackageInfo();
                $('#editBusinessPackageModel').modal('show');
                vc.copyObject(_params, vc.component.editBusinessPackageInfo);
            });
        },
        methods: {
            editBusinessPackageValidate: function() {
                return vc.validate.validate({
                    editBusinessPackageInfo: vc.component.editBusinessPackageInfo
                }, {
                    'editBusinessPackageInfo.name': [{
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
                    'editBusinessPackageInfo.id': [{
                        limit: "required",
                        param: "",
                        errInfo: "ID不能为空"
                    }]

                });
            },
            editBusinessPackage: function() {
                if (!vc.component.editBusinessPackageValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                var param = new FormData();
                param.append("uploadFile", vc.component.editBusinessPackageInfo.excelEditTemplate);
                param.append('name', vc.component.editBusinessPackageInfo.name);
                param.append('id', vc.component.editBusinessPackageInfo.id);
                param.append('type', vc.component.addBusinessPackageInfo.type);
                param.append('filename', vc.component.addBusinessPackageInfo.filename);

                vc.http.apiPost(
                    '/soft/updateBusinessPackages',
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
                            $('#editBusinessPackageModel').modal('hide');
                            vc.component.clearEditBusinessPackageInfo();
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
            getEditExcelTemplate: function(e) {
                //console.log("getExcelTemplate 开始调用")
                vc.component.editBusinessPackageInfo.excelEditTemplate = e.target.files[0];
            },
            refreshEditBusinessPackageInfo: function() {
                vc.component.editBusinessPackageInfo = {
                    id: '',
                    name: '',
                    varsion: '',
                    excelEditTemplate: '',
                    type: '1001',
                    filename: ''
                }
            }
        }
    });

})(window.vc, window.vc.component);