(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addAppVarGroupInfo: {
                avgId: '',
                avgName: '',
                avgDesc: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addAppVarGroup', 'openAddAppVarGroupModal', function () {
                
                $('#addAppVarGroupModel').modal('show');
            });
        },
        methods: {
            addAppVarGroupValidate() {
                return vc.validate.validate({
                    addAppVarGroupInfo: vc.component.addAppVarGroupInfo
                }, {
                    'addAppVarGroupInfo.avgName': [
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
                    'addAppVarGroupInfo.avgDesc': [
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




                });
            },
            saveAppVarGroupInfo: function () {
                if (!vc.component.addAppVarGroupValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addAppVarGroupInfo);
                    $('#addAppVarGroupModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/appService/saveAppVarGroup',
                    JSON.stringify(vc.component.addAppVarGroupInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addAppVarGroupModel').modal('hide');
                            vc.component.clearAddAppVarGroupInfo();
                            vc.emit('appVarGroupManage', 'listAppVarGroup', {});

                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            clearAddAppVarGroupInfo: function () {
                vc.component.addAppVarGroupInfo = {
                    avgName: '',
                    avgDesc: '',

                };
            },
            
        }
    });

})(window.vc);
