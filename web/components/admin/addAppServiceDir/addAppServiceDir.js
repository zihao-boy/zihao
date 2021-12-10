(function(vc) {

    vc.extends({
        data: {
            addAppServiceDirInfo: {
                asId: '',
                srcDir: '',
                targetDir: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('addAppServiceDir', 'openAddAppServiceDirModal', function(_param) {
                $that.addAppServiceDirInfo.asId = _param.asId;
                $('#addAppServiceDirModel').modal('show');
            });
        },
        methods: {
            addAppServiceDirValidate() {
                return vc.validate.validate({
                    addAppServiceDirInfo: vc.component.addAppServiceDirInfo
                }, {
                    'addAppServiceDirInfo.srcDir': [{
                        limit: "required",
                        param: "",
                        errInfo: "宿主机目录不能为空"
                    }],
                    'addAppServiceDirInfo.targetDir': [{
                        limit: "required",
                        param: "",
                        errInfo: "容器目录不能为空"
                    }]
                });
            },
            saveAppServiceDirInfo: function() {
                if (!vc.component.addAppServiceDirValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appService/saveAppServiceDir',
                    JSON.stringify(vc.component.addAppServiceDirInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppServiceDirModel').modal('hide');
                            vc.emit('appServiceControlDir', 'switch', $that.addAppServiceDirInfo);
                            vc.component.clearaddAppServiceDirInfo();

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearaddAppServiceDirInfo: function() {
                vc.component.addAppServiceDirInfo = {
                    asId: '',
                    srcDir: '',
                    targetDir: ''
                };
            }
        }
    });

})(window.vc);