(function (vc, vm) {

    vc.extends({
        data: {
            editServiceSqlInfo: {
                sqlId: '',
                sqlCode: '',
                remark: '',
                sqlText: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editServiceSql', 'openEditServiceSqlModal', function (_params) {
                vc.component.refreshEditServiceSqlInfo();
                $('#editServiceSqlModel').modal('show');
                vc.copyObject(_params, vc.component.editServiceSqlInfo);
                vc.component.editServiceSqlInfo.communityId = vc.getCurrentCommunity().communityId;
            });
        },
        methods: {
            editServiceSqlValidate: function () {
                return vc.validate.validate({
                    editServiceSqlInfo: vc.component.editServiceSqlInfo
                }, {
                    'editServiceSqlInfo.sqlCode': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "sql编码不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "sql编码格式错误"
                        },
                    ],
                    'editServiceSqlInfo.remark': [
                        {
                            limit: "maxLength",
                            param: "512",
                            errInfo: "备注太长"
                        },
                    ],
                    'editServiceSqlInfo.sqlText': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "sql语句不能为空"
                        }
                    ],
                    'editServiceSqlInfo.sqlId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "Sql ID不能为空"
                        }]

                });
            },
            editServiceSql: function () {
                if (!vc.component.editServiceSqlValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/system/updateServiceSql',
                    JSON.stringify(vc.component.editServiceSqlInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editServiceSqlModel').modal('hide');
                            vc.emit('serviceSqlManage', 'listServiceSql', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditServiceSqlInfo: function () {
                vc.component.editServiceSqlInfo = {
                    sqlId: '',
                    sqlCode: '',
                    remark: '',
                    sqlText: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
