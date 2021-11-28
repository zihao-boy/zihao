(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostMonitorsInfo: {
                hostId: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            //切换 至费用页面
            vc.on('hostMonitors', 'switch', function (_param) {
                if (_param.hostId == '') {
                    return;
                }
                vc.copyObject(_param, $that.hostMonitorsInfo)
                $that._hostsMonitorsData(DEFAULT_PAGE, DEFAULT_ROWS);
            });
        },
        methods: {
            _hostsMonitorsData: function () {

                let param = {
                    params: {
                        hostId:$that.hostMonitorsInfo.hostId
                    }
                }

                //发送get请求
                vc.http.apiGet('/monitor/getMonitorHostLog',
                    param,
                    function (json, res) {
                        let _json = JSON.parse(json);
                        let _data = _json.data;
                        let _xAxis = [];
                        let _cpuData = [];
                        let _memData = [];
                        let _diskData = [];
                        _data.forEach(item => {
                            _xAxis.push(item.createTime + "时");
                            _cpuData.push(item.cpuRate * 100);
                            _memData.push(item.memRate * 100);
                            _diskData.push(item.diskRate * 100);
                        });
                        let _dom = document.getElementById('hostMonitors');
                        $that._initEcharts(_dom, _xAxis, _cpuData, _memData, _diskData);
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );

            },
            _initEcharts: function (dom, _xAxis, _cpuData, _memData, _diskData) {
                //let dom = document.getElementById("box2");
                let myChart = echarts.init(dom);
                let option = null;
                option = {
                    tooltip: {
                        trigger: 'axis'
                    },
                    legend: {
                        data: ['cpu使用率', '内存使用率', '磁盘使用率']
                    },
                    grid: {
                        left: '3%',
                        right: '4%',
                        bottom: '3%',
                        containLabel: true
                    },
                    toolbox: {
                        feature: {
                            saveAsImage: {}
                        }
                    },
                    xAxis: {
                        type: 'category',
                        boundaryGap: false,
                        data: _xAxis
                    },
                    yAxis: {
                        type: 'value'
                    },
                    series: [
                        {
                            name: 'cpu使用率',
                            type: 'line',
                            data: _cpuData
                        },
                        {
                            name: '内存使用率',
                            type: 'line',
                            data: _memData
                        },
                        {
                            name: '磁盘使用率',
                            type: 'line',
                            data: _diskData
                        }
                    ]
                };;

                if (option && typeof option === "object") {
                    myChart.setOption(option, true);
                }

            },

        }

    });
})(window.vc);
