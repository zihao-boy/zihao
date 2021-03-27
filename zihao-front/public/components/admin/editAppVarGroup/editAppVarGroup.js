(function (vc, vm) {

    vc.extends({
        data: {
            editAppVarGroupInfo: {
                avgId: '',
                avgName: '',
                avgDesc: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editAppVarGroup', 'openEditAppVarGroupModal', function (_params) {
                vc.component.refreshEditAppVarGroupInfo();
                $('#editAppVarGroupModel').modal('show');
                vc.copyObject(_params, vc.component.editAppVarGroupInfo);
                vc.component.editAppVarGroupInfo.communityId = vc.getCurrentCommunity().communityId;
            });
        },
        methods: {
            editAppVarGroupValidate: function () {
                return vc.validate.validate({
                    editAppVarGroupInfo: vc.component.editAppVarGroupInfo
                }, {
                    'editAppVarGroupInfo.avgName': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "组名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "组名称太长"
                        },
                    ],
                    'editAppVarGroupInfo.avgDesc': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "组描述不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "512",
                            errInfo: "描述太长"
                        },
                    ],
                    'editAppVarGroupInfo.avgId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "组编码不能为空"
                        }]

                });
            },
            editAppVarGroup: function () {
                if (!vc.component.editAppVarGroupValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/appService/updateAppVarGroup',
                    JSON.stringify(vc.component.editAppVarGroupInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editAppVarGroupModel').modal('hide');
                            vc.emit('appVarGroupManage', 'listAppVarGroup', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditAppVarGroupInfo: function () {
                vc.component.editAppVarGroupInfo = {
                    avgId: '',
                    avgName: '',
                    avgDesc: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
