(function(vc, vm) {

    vc.extends({
        data: {
            editFtpInfo: {
                ftpId: '',
                name: '',
                ip: '',
                port: '',
                username: '',
                passwd: '',
                path: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editFtp', 'openEditFtpModal', function(_params) {
                vc.component.refreshEditFtpInfo();
                $('#editFtpModel').modal('show');
                vc.copyObject(_params, vc.component.editFtpInfo);
            });
        },
        methods: {
            editFtpValidate: function() {
                return vc.validate.validate({
                    editFtpInfo: vc.component.editFtpInfo
                }, {
                    'editFtpInfo.name': [{
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
                    'editFtpInfo.ip': [{
                            limit: "required",
                            param: "",
                            errInfo: "IP不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "IP不能超过128"
                        },
                    ],
                    'editFtpInfo.port': [{
                            limit: "required",
                            param: "",
                            errInfo: "端口不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "端口不能超过64"
                        },
                    ],
                    'editFtpInfo.username': [{
                            limit: "required",
                            param: "",
                            errInfo: "用户名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "用户名不能超过64"
                        },
                    ],
                    'editFtpInfo.passwd': [{
                            limit: "required",
                            param: "",
                            errInfo: "密码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "密码不能超过64"
                        },
                    ],
                    'editFtpInfo.path': [{
                            limit: "required",
                            param: "",
                            errInfo: "路径'不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "路径'不能超过128"
                        },
                    ],
                    'editFtpInfo.ftpId': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editFtp: function() {
                if (!vc.component.editFtpValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/resources/updateFtp',
                    JSON.stringify(vc.component.editFtpInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editFtpModel').modal('hide');
                            vc.emit('ftpManage', 'listFtp', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditFtpInfo: function() {
                vc.component.editFtpInfo = {
                    ftpId: '',
                    name: '',
                    ip: '',
                    port: '',
                    username: '',
                    passwd: '',
                    path: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);