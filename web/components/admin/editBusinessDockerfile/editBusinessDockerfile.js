(function(vc, vm) {

    vc.extends({
        data: {
            editBusinessDockerfileInfo: {
                id: '',
                name: '',
                dockerfile: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editBusinessDockerfile', 'openEditBusinessDockerfileModal', function(_params) {
                vc.component.refreshEditBusinessDockerfileInfo();
                $('#editBusinessDockerfileModel').modal('show');
                vc.copyObject(_params, vc.component.editBusinessDockerfileInfo);
            });
        },
        methods: {
            editBusinessDockerfileValidate: function() {
                return vc.validate.validate({
                    editBusinessDockerfileInfo: vc.component.editBusinessDockerfileInfo
                }, {
                    'editBusinessDockerfileInfo.name': [{
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
                    'editBusinessDockerfileInfo.dockerfile': [{
                        limit: "required",
                        param: "",
                        errInfo: "内容不能为空"
                    }],
                    'editBusinessDockerfileInfo.id': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editBusinessDockerfile: function() {
                if (!vc.component.editBusinessDockerfileValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/soft/updateBusinessDockerfile',
                    JSON.stringify(vc.component.editBusinessDockerfileInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editBusinessDockerfileModel').modal('hide');
                            vc.emit('businessDockerfileManage', 'listBusinessDockerfile', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditBusinessDockerfileInfo: function() {
                vc.component.editBusinessDockerfileInfo = {
                    id: '',
                    name: '',
                    dockerfile: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);