(function (vc, vm) {

    vc.extends({
        data: {
            editMonitorHostInfo: {
                mhId: '',
                hostId: '',
                monDisk: '',
                cpuThreshold: '',
                memThreshold: '',
                diskThreshold: '',

            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            vc.on('editMonitorHost', 'openEditMonitorHostModal', function (_params) {
                vc.component.refreshEditMonitorHostInfo();
                $('#editMonitorHostModel').modal('show');
                vc.copyObject(_params, vc.component.editMonitorHostInfo);
               
            });
        },
        methods: {
            editMonitorHostValidate: function () {
                return vc.validate.validate({
                    editMonitorHostInfo: vc.component.editMonitorHostInfo
                }, {
                    'editMonitorHostInfo.hostId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "主机不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "主机错误"
                        },
                    ],
                    'editMonitorHostInfo.monDisk': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "监控目录不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "128",
                            errInfo: "监控目录太长"
                        },
                    ],
                    'editMonitorHostInfo.cpuThreshold': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "cpu阈值不能为空"
                        },
                        {
                            limit: "money",
                            param: "",
                            errInfo: "cpu阈值错误,正确如0.80"
                        },
                    ],
                    'editMonitorHostInfo.memThreshold': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "内存阈值不能为空"
                        },
                        {
                            limit: "money",
                            param: "",
                            errInfo: "内存阈值错误,正确如0.80"
                        },
                    ],
                    'editMonitorHostInfo.diskThreshold': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "磁盘阈值不能为空"
                        },
                        {
                            limit: "money",
                            param: "",
                            errInfo: "磁盘阈值错误,正确如0.80"
                        },
                    ],
                    'editMonitorHostInfo.mhId': [
                        {
                            limit: "required",
                            param: "",
                            errInfo: "ID不能为空"
                        }]

                });
            },
            editMonitorHost: function () {
                if (!vc.component.editMonitorHostValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/monitor/updateMonitorHost',
                    JSON.stringify(vc.component.editMonitorHostInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editMonitorHostModel').modal('hide');
                            vc.emit('monitorHostManage', 'listMonitorHost', {});
                            return;
                        }
                        vc.message(_json.msg);
                    },
                    function (errInfo, error) {
                        console.log('请求失败处理');

                        vc.message(errInfo);
                    });
            },
            refreshEditMonitorHostInfo: function () {
                vc.component.editMonitorHostInfo = {
                    mhId: '',
                    hostId: '',
                    monDisk: '',
                    cpuThreshold: '',
                    memThreshold: '',
                    diskThreshold: '',

                }
            }
        }
    });

})(window.vc, window.vc.component);
