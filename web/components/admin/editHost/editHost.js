(function(vc, vm) {

    vc.extends({
        data: {
            editHostInfo: {
                hostId: '',
                groupId: '',
                name: '',
                ip: '',
                username: '',
                passwd: '',
                hostGroups: [],
                osName:''

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editHost', 'openEditHostModal', function(_params) {
                vc.component.refreshEditHostInfo();
                $that.loadEditHostHostGroup();
                $('#editHostModel').modal('show');
                vc.copyObject(_params, vc.component.editHostInfo);
            });
        },
        methods: {
            editHostValidate: function() {
                return vc.validate.validate({
                    editHostInfo: vc.component.editHostInfo
                }, {
                    'editHostInfo.groupId': [{
                            limit: "required",
                            param: "",
                            errInfo: "主机组不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "主机组格式错误"
                        },
                    ],
                    'editHostInfo.name': [{
                            limit: "required",
                            param: "",
                            errInfo: "主机名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "主机名称太长"
                        },
                    ],
                    'editHostInfo.ip': [{
                            limit: "required",
                            param: "",
                            errInfo: "主机IP不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "主机IP太长"
                        },
                    ],
                    'editHostInfo.username': [{
                            limit: "required",
                            param: "",
                            errInfo: "用户名不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "用户名太长"
                        },
                    ],
                    'editHostInfo.passwd': [{
                            limit: "required",
                            param: "",
                            errInfo: "密码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "密码太长"
                        },
                    ],
                    'editHostInfo.hostId': [{
                        limit: "required",
                        param: "",
                        errInfo: "主机ID不能为空"
                    }]

                });
            },
            editHost: function() {
                if (!vc.component.editHostValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/host/updateHost',
                    JSON.stringify(vc.component.editHostInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editHostModel').modal('hide');
                            vc.emit('hostManage', 'listHost', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditHostInfo: function() {
                vc.component.editHostInfo = {
                    hostId: '',
                    groupId: '',
                    name: '',
                    ip: '',
                    username: '',
                    passwd: '',
                    hostGroups: [],
                    osName:''
                }
            },
            loadEditHostHostGroup: function() {
                var param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHostGroup',
                    param,
                    function(json, res) {
                        var _hostGroupManageInfo = JSON.parse(json);
                        vc.component.editHostInfo.hostGroups = _hostGroupManageInfo.data;

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            }
        }
    });

})(window.vc, window.vc.component);