/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            monitorHostGroupManageAnalysisInfo: {
                hostGroups: [],
                hosts: [],
                total: 0,
                records: 1,
                hasCondition: true,
                groupId: '',
                conditions: {
                    hostId: '0638e278-484c-4e49-b59f-9ed73ab418ed',
                }
            }
        },
        _initMethod: function () {
            let _hostId = vc.getParam('hostId');
            if(vc.notNull(_hostId)){
                $that.monitorHostGroupManageAnalysisInfo.hasCondition = false;
                $that.monitorHostGroupManageAnalysisInfo.conditions.hostId = _hostId;
                $that._hostsData();
            }else{
                $that._listHostGroups();
            }
        },
        _initEvent: function () {

            vc.on('monitorHostGroupManage', 'listMonitorHostGroup', function (_param) {
                vc.component._listMonitorHostGroups(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listMonitorHostGroups(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _hostsData: function () {

                let param = {
                    params: $that.monitorHostGroupManageAnalysisInfo.conditions
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
                            _cpuData.push(item.cpuRate);
                            _memData.push(item.memRate * 100);
                            _diskData.push(item.diskRate * 100);
                        });
                        let _dom = document.getElementById('hostAnalysis');
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
            _listHostGroups: function () {
                var param = {
                    params: {
                        page: 1,
                        row: 50
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHostGroup',
                    param,
                    function (json, res) {
                        var _hostGroupManageInfo = JSON.parse(json);
                       $that.monitorHostGroupManageAnalysisInfo.hostGroups = _hostGroupManageInfo.data;
                        if(_hostGroupManageInfo.data.length > 0){
                            $that.monitorHostGroupManageAnalysisInfo.groupId = _hostGroupManageInfo.data[0].groupId;
                            $that._listHosts();
                        }else{
                            $that.monitorHostGroupManageAnalysisInfo.hosts=[]
                        }
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },

            _listHosts: function () {

                var param = {
                    params: {
                        groupId:$that.monitorHostGroupManageAnalysisInfo.groupId,
                        row:100,
                        page:1
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHosts',
                    param,
                    function (json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        $that.monitorHostGroupManageAnalysisInfo.hosts = _hostManageInfo.data;
                        if(_hostManageInfo.data.length > 0){
                            $that.monitorHostGroupManageAnalysisInfo.conditions.hostId = _hostManageInfo.data[0].hostId;
                            $that._hostsData();
                        }
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _goBack:function(){
                vc.goBack();
            }
        }
    });
})(window.vc);
