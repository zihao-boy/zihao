(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addServiceSqlInfo: {
                sqlId: '',
                sqlCode: '',
                remark: '',
                sqlText: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addServiceSql', 'openAddServiceSqlModal', function () {
                $('#addServiceSqlModel').modal('show');
            });
        },
        methods: {
            addServiceSqlValidate() {
                return vc.validate.validate({
                    addServiceSqlInfo: vc.component.addServiceSqlInfo
                }, {
                    'addServiceSqlInfo.sqlCode': [
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
                    'addServiceSqlInfo.remark': [
                        {
                            limit: "maxLength",
                            param: "512",
                            errInfo: "备注太长"
                        },
                    ],
                    'addServiceSqlInfo.sqlText': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "sql语句不能为空"
                        }
                    ],

                });
            },
            saveServiceSqlInfo: function () {
                if (!vc.component.addServiceSqlValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                vc.component.addServiceSqlInfo.communityId = vc.getCurrentCommunity().communityId;
                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addServiceSqlInfo);
                    $('#addServiceSqlModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    'serviceSql.saveServiceSql',
                    JSON.stringify(vc.component.addServiceSqlInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addServiceSqlModel').modal('hide');
                            vc.component.clearAddServiceSqlInfo();
                            vc.emit('serviceSqlManage', 'listServiceSql', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddServiceSqlInfo: function () {
                vc.component.addServiceSqlInfo = {
                    sqlCode: '',
                    remark: '',
                    sqlText: '',

                };
            }
        }
    });

})(window.vc);
