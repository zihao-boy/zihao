(function (vc) {

    vc.extends({
        propTypes: {
            callBackListener: vc.propTypes.string, //父组件名称
            callBackFunction: vc.propTypes.string //父组件监听方法
        },
        data: {
            addMonitorHostInfo: {
                mhId: '',
                hostId: '',
                monDisk: '',
                cpuThreshold: '',
                memThreshold: '',
                diskThreshold: '',
                hosts:[],
                mhgId:''
            }
        },
        _initMethod: function () {
            $that._listHosts()
        },
        _initEvent: function () {
            vc.on('addMonitorHost', 'openAddMonitorHostModal', function (_data) {
                $that.addMonitorHostInfo.mhgId = _data.mhgId
                $('#addMonitorHostModel').modal('show');
            });
        },
        methods: {
            addMonitorHostValidate() {
                return vc.validate.validate({
                    addMonitorHostInfo: vc.component.addMonitorHostInfo
                }, {
                    'addMonitorHostInfo.hostId': [
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
                    'addMonitorHostInfo.monDisk': [
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
                    'addMonitorHostInfo.cpuThreshold': [
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
                    'addMonitorHostInfo.memThreshold': [
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
                    'addMonitorHostInfo.diskThreshold': [
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




                });
            },
            saveMonitorHostInfo: function () {
                if (!vc.component.addMonitorHostValidate()) {
                    vc.toast(vc.validate.errInfo);

                    return;
                }

                //不提交数据将数据 回调给侦听处理
                if (vc.notNull($props.callBackListener)) {
                    vc.emit($props.callBackListener, $props.callBackFunction, vc.component.addMonitorHostInfo);
                    $('#addMonitorHostModel').modal('hide');
                    return;
                }

                vc.http.apiPost(
                    '/monitor/saveMonitorHost',
                    JSON.stringify(vc.component.addMonitorHostInfo),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#addMonitorHostModel').modal('hide');
                            vc.component.clearAddMonitorHostInfo();
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
            clearAddMonitorHostInfo: function () {
                let _hosts = $that.addMonitorHostInfo.hosts
                vc.component.addMonitorHostInfo = {
                    hostId: '',
                    monDisk: '',
                    cpuThreshold: '',
                    memThreshold: '',
                    diskThreshold: '',
                    hosts:_hosts,
                    mhgId:''
                };
            },
            _listHosts: function (_page, _rows) {
                var param = {
                    params: {
                        page:1,
                        row:100
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function (json, res) {
                        var _hostManageInfo = JSON.parse(json);
                        vc.component.addMonitorHostInfo.hosts = _hostManageInfo.data;
                       
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
        }
    });

})(window.vc);
