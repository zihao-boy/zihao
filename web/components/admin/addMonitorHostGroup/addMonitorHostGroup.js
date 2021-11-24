(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addMonitorHostGroupInfo: {
                mhgId: '',
                name: '',
                monCron: '',
                noticeType: '',
                remark: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('addMonitorHostGroup', 'openAddMonitorHostGroupModal', function () {
                $('#addMonitorHostGroupModel').modal('show');
            });
        },
        methods: {
            addMonitorHostGroupValidate() {
                return vc.validate.validate({
                    addMonitorHostGroupInfo: vc.component.addMonitorHostGroupInfo
                }, {
                    'addMonitorHostGroupInfo.name': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "名称不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "名称太长"
                        },
                    ],
                    'addMonitorHostGroupInfo.monCron': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "监控频率不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "监控频率太长"
                        },
                    ],
                    'addMonitorHostGroupInfo.noticeType': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "告警类型不能为空"
                        },
                        {
                            limit: "num",
                            param: "",
                            errInfo: "告警类型错误"
                        },
                    ],
                    'addMonitorHostGroupInfo.remark': [
                        {
                            limit: "maxLength",
                            param: "1000",
                            errInfo: "备注太长"
                        },
                    ],




                });
            },
            saveMonitorHostGroupInfo: function () {
                if (!vc.component.addMonitorHostGroupValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addMonitorHostGroupInfo);
                    $('#addMonitorHostGroupModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/monitor/saveMonitorHostGroup',
                    JSON.stringify(vc.component.addMonitorHostGroupInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addMonitorHostGroupModel').modal('hide');
                            vc.component.clearAddMonitorHostGroupInfo();
                            vc.emit('monitorHostGroupManage', 'listMonitorHostGroup', {});

                            return;
                        }
                        vc.toast(_json.msg);

                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);

                    });
            },
            clearAddMonitorHostGroupInfo: function () {
                vc.component.addMonitorHostGroupInfo = {
                    name: '',
                    monCron: '',
                    noticeType: '',
                    remark: '',

                };
            }
        }
    });

})(window.vc);
