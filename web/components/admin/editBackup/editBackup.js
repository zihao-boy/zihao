(function(vc, vm) {

    vc.extends({
        data: {
            editBackupInfo: {
                id: '',
                name: '',
                execTime: '',
                typeCd: '',
                srcId: '',
                srcObject: '',
                targetId: '',

            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('editBackup', 'openEditBackupModal', function(_params) {
                vc.component.refreshEditBackupInfo();
                $('#editBackupModel').modal('show');
                vc.copyObject(_params, vc.component.editBackupInfo);
            });
        },
        methods: {
            editBackupValidate: function() {
                return vc.validate.validate({
                    editBackupInfo: vc.component.editBackupInfo
                }, {
                    'editBackupInfo.name': [{
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
                    'editBackupInfo.execTime': [{
                            limit: "required",
                            param: "",
                            errInfo: "执行时间不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "执行时间不能超过64"
                        },
                    ],
                    'editBackupInfo.typeCd': [{
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
                    'editBackupInfo.srcId': [{
                            limit: "required",
                            param: "",
                            errInfo: "备份编号不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "备份编号不能超过64"
                        },
                    ],
                    'editBackupInfo.srcObject': [{
                            limit: "required",
                            param: "",
                            errInfo: "备份对象不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "1024",
                            errInfo: "备份对象不能超过1024"
                        },
                    ],
                    'editBackupInfo.targetId': [{
                            limit: "required",
                            param: "",
                            errInfo: "目标不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "目标不能超过64"
                        },
                    ],
                    'editBackupInfo.id': [{
                        limit: "required",
                        param: "",
                        errInfo: "编号不能为空"
                    }]

                });
            },
            editBackup: function() {
                if (!vc.component.editBackupValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/resources/updateBackUp',
                    JSON.stringify(vc.component.editBackupInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editBackupModel').modal('hide');
                            vc.emit('backupManage', 'listBackup', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditBackupInfo: function() {
                vc.component.editBackupInfo = {
                    id: '',
                    name: '',
                    execTime: '',
                    typeCd: '',
                    srcId: '',
                    srcObject: '',
                    targetId: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);