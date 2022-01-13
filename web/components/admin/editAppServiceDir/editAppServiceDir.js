(function(vc) {

    vc.extends({
        data: {
            editAppServiceDirInfo: {
                asId: '',
                srcDir: '',
                targetDir: '',
                dirId: ''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editAppServiceDir', 'openEditAppServiceDirModal', function(_param) {
                vc.copyObject(_param, $that.editAppServiceDirInfo);
                $('#editAppServiceDirModel').modal('show');
            });
        },
        methods: {
            editAppServiceDirValidate() {
                return vc.validate.validate({
                    editAppServiceDirInfo: vc.component.editAppServiceDirInfo
                }, {
                    'editAppServiceDirInfo.srcDir': [{
                        limit: "required",
                        param: "",
                        errInfo: "宿主机目录不能为空"
                    }],
                    'editAppServiceDirInfo.targetDir': [{
                        limit: "required",
                        param: "",
                        errInfo: "容器目录不能为空"
                    }]
                });
            },
            updateAppServiceDirInfo: function() {
                if (!vc.component.editAppServiceDirValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }
                vc.http.apiPost(
                    '/appService/updateAppServiceDir',
                    JSON.stringify(vc.component.editAppServiceDirInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppServiceDirModel').modal('hide');
                            vc.emit('appServiceControlDir', 'switch', $that.editAppServiceDirInfo);
                            vc.component.cleareditAppServiceDirInfo();

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            cleareditAppServiceDirInfo: function() {
                vc.component.editAppServiceDirInfo = {
                    asId: '',
                    srcDir: '',
                    targetDir: '',
                    dirId: ''
                };
            }
        }
    });

})(window.vc);