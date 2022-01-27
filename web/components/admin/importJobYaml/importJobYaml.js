(function(vc, vm) {
    vc.extends({
        data: {
            importJobYamlInfo: {
                excelTemplate: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('importJobYaml', 'openImportJobYamlModal', function(_params) {
                $('#importJobYamlModel').modal('show');
            });
        },
        methods: {
            importJobYaml: function() {

                let param = new FormData();
                param.append("uploadFile", vc.component.importJobYamlInfo.excelTemplate);

                vc.http.apiPost(
                    '/appVersion/importJobBuildYaml',
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
                            $('#importJobYamlModel').modal('hide');
                            vc.component.refreshImportJobYamlInfo();
                            vc.emit('appVersionJobManage', 'listAppVersionJob', {});
                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                        vc.toast(errInfo);
                    });
            },
            refreshImportJobYamlInfo: function() {
                vc.component.importJobYamlInfo = {
                    excelTemplate: ''
                }
            },
            getExcelTemplate: function(e) {
                //console.log("getExcelTemplate 开始调用")
                vc.component.importJobYamlInfo.excelTemplate = e.target.files[0];
            },
        }
    });

})(window.vc, window.vc.component);