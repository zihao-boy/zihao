(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            hostResourcesInfo: {
                hostId: ''
            }
        },
        _initMethod: function () {

        },
        _initEvent: function () {
            //切换 至费用页面
            vc.on('hostResources', 'switch', function (_param) {
                if (_param.hostId == '') {
                    return;
                }
                vc.copyObject(_param, $that.hostResourcesInfo)
                $that._hostResourcesData(DEFAULT_PAGE, DEFAULT_ROWS);
            });
        },
        methods: {
            _hostResourcesData: function () {

                let param = {
                    params: {
                        hostId:$that.hostResourcesInfo.hostId
                    }
                }

                //发送get请求
                vc.http.apiGet('/host/getHostResource',
                    param,
                    function (json, res) {
                        let _json = JSON.parse(json);
                        let _data = _json.data;

                        //cpu
                        let _dom = document.getElementById('hostCpu');
                        let _newdata = [
                            { value: 100-_data.cpuRate, name: '空闲' },
                            { value: _data.cpuRate, name: '已使用' }
                        ];
                        $that._initHostResourceCharts(_dom, 'cpu使用率', _newdata);

                        //内存
                         _dom = document.getElementById('hostMem');
                         _newdata = [
                            { value: _data.memTotal-_data.memUsed, name: '空闲' },
                            { value: _data.memUsed, name: '已使用' }
                        ];
                        $that._initHostResourceCharts(_dom, '内存使用率', _newdata);

                         //磁盘使用率
                         _dom = document.getElementById('hostDisk');
                         _newdata = [
                            { value: _data.diskTotal-_data.diskUsed, name: '空闲' },
                            { value: _data.diskUsed, name: '已使用' }
                        ];
                        $that._initHostResourceCharts(_dom, '磁盘使用率', _newdata);
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );

            },
            _initHostResourceCharts: function (dom, _title, _data) {

                let myChart = echarts.init(dom);
                let option = null;
                option = {
                    title: {
                        text: _title,
                        left: 'center'
                    },
                    textStyle: {//图例文字的样式
                        fontSize: 12
                    },
                    tooltip: {
                        trigger: 'item',
                        formatter: '{a} <br/>{b} : {c} ({d}%)'
                    },
                    color: ['#66CDAA', '#FFDAB9'],
                    series: [
                        {
                            name: _title,
                            type: 'pie',
                            radius: '75%',
                            center: ['50%', '50%'],
                            data: _data,
                            emphasis: {
                                itemStyle: {
                                    shadowBlur: 10,
                                    shadowOffsetX: 0,
                                    shadowColor: 'rgba(0, 0, 0, 0.5)'
                                }
                            },
                            label: {
                                normal: {
                                    show: false
                                }
                            }
                        }
                    ]
                };
                if (option && typeof option === "object") {
                    myChart.setOption(option, true);
                }
            },

        }

    });
})(window.vc);
