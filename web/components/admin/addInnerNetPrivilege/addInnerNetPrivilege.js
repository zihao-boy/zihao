(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addInnerNetPrivilegeInfo: {
                pId: '',
                srcUserId: '',
                targetUserId: '',
                targetPort: '',
                innerNetUsers:[]

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addInnerNetPrivilege', 'openAddInnerNetPrivilegeModal', function () {
                $that._listAddInnerNetUsers();
                $('#addInnerNetPrivilegeModel').modal('show');
            });
        },
        methods: {
            addInnerNetPrivilegeValidate() {
                return vc.validate.validate({
                    addInnerNetPrivilegeInfo: vc.component.addInnerNetPrivilegeInfo
                }, {
                    'addInnerNetPrivilegeInfo.srcUserId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "用户不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "用户不能超过64"
                        },
                    ],
                    'addInnerNetPrivilegeInfo.targetUserId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "主机不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "主机不能超过64"
                        },
                    ],
                    'addInnerNetPrivilegeInfo.targetPort': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "端口'不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "端口'不能超过64"
                        },
                    ],




                });
            },
            _listAddInnerNetUsers: function () {

                var param = {
                    params: {
                        page:1,
                        row:1000
                    }
                };

                //发送get请求
                vc.http.apiGet('/innerNet/getInnerNetUser',
                    param,
                    function (json, res) {
                        var _innerNetUserManageInfo = JSON.parse(json);
                        vc.component.addInnerNetPrivilegeInfo.innerNetUsers = _innerNetUserManageInfo.data;
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            saveInnerNetPrivilegeInfo: function () {
                if (!vc.component.addInnerNetPrivilegeValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addInnerNetPrivilegeInfo);
                    $('#addInnerNetPrivilegeModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/innerNet/saveInnerNetPrivilege',
                    JSON.stringify(vc.component.addInnerNetPrivilegeInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addInnerNetPrivilegeModel').modal('hide');
                            vc.component.clearAddInnerNetPrivilegeInfo();
                            vc.emit('innerNetPrivilegeManage', 'listInnerNetPrivilege', {});

                            return;
                        }
                        vc.message(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);

                    });
            },
            clearAddInnerNetPrivilegeInfo: function () {
                vc.component.addInnerNetPrivilegeInfo = {
                    srcUserId: '',
                    targetUserId: '',
                    targetPort: '',
                    innerNetUsers:[]

                };
            }
        }
    });

})(window.vc);
