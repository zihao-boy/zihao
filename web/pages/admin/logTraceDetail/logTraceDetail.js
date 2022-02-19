/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            logTraceDetailInfo: {
                traces: [],
                traceId: '',
                spanId: '',
                serviceName: '',
                name: '',
                reqHeader: '',
                reqParam: '',
                resParam: '',
                resHeader: '',
                errInfo: '',
                annos: []
            }
        },
        _initMethod: function() {
            $that.logTraceDetailInfo.traceId = vc.getParam('traceId');
            vc.component._listLogTraceDetails(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

        },
        methods: {
            _listLogTraceDetails: function(_page, _rows) {
                let param = {
                    params: {
                        traceId: $that.logTraceDetailInfo.traceId,
                        page: 1,
                        row: 50
                    }
                };

                //发送get请求
                vc.http.apiGet('/monitor/getLogTraceDetail',
                    param,
                    function(json, res) {
                        let _logTraceInfo = JSON.parse(json);
                        vc.component.logTraceDetailInfo.total = _logTraceInfo.total;
                        vc.component.logTraceDetailInfo.records = _logTraceInfo.records;
                        vc.component.logTraceDetailInfo.traces = _logTraceInfo.data.sort($that.compare);
                        if (_logTraceInfo.data.length < 1) {
                            return;
                        }
                        $that._filterData();
                        $that._loadReqInfo(_logTraceInfo.data[0].id);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _filterData: function() {
                let _nodes = [];
                let _edges = [];

                let _traces = $that.logTraceDetailInfo.traces;

                let _serviceName = "";

                let _edgesName = "";
                let _annotations = 0;
                _traces.forEach((trace) => {
                    //_annotations = trace.annotations;

                    _serviceName = "服务名：" + trace.serviceName +
                        "\n接口名：" + trace.name +
                        "\n主机信息：" + trace.ipv4 + ":" + trace.port +
                        "\n业务耗时：" + trace.duration +
                        "\n调用时间：" + trace.timestamp ;
                    _nodes.push({
                        "id": trace.id,
                        "label": _serviceName,
                        "shape": "rect"
                    });

                    if (trace.parentSpanId == "0") {
                        return;
                    }

                    _edgesName = "网络耗时：" + $that._computeNetTime(trace)

                    _edges.push({
                        "label": _edgesName,
                        "source": trace.parentSpanId,
                        "target": trace.id
                    })
                })

                $that._initTraceDetail({
                    edges: _edges,
                    nodes: _nodes
                })
            },

            // cs - ss
            _computeNetTime: function(_trace) {
                // 寻找父节点
                let _traces = $that.logTraceDetailInfo.traces;
                let _annotations = {};
                let _ssTime = 0;
                let _csTime = 0;
                _traces.forEach((item) => {
                    if (item.id != _trace.parentSpanId) {
                        return;
                    }

                    //_annotations = item.annotations;
                    _annotations = item.annotations.sort($that.compare);
                    _annotations.forEach((anno) => {
                        if (anno.value == 'ss' &&  _ssTime == 0) {
                            _ssTime = anno.timestamp;
                        }
                    })
                });
                _annotations = _trace.annotations;
                _annotations.forEach((anno) => {
                    if (anno.value == 'cs') {
                        _csTime = anno.timestamp;
                    }
                })

                return _csTime - _ssTime;

            },
            _initTraceDetail: function(_dataset) {
                let g = new dagreD3.graphlib.Graph();
                //设置图
                g.setGraph({
                    rankdir: 'TB'
                });
                _dataset.nodes.forEach(item => {
                    g.setNode(item.id, {
                        //节点标签
                        label: item.label,
                        //节点形状
                        shape: item.shape,
                        //节点样式
                        style: "fill:#fff;stroke:#000"
                    })
                })
                _dataset.edges.forEach(item => {
                        g.setEdge(item.source, item.target, {
                            //边标签
                            label: item.label,
                            //边样式
                            style: "fill:#fff;stroke:#333;stroke-width:1.5px"
                        })
                    })
                    // debugger
                    // 创建渲染器
                let render = new dagreD3.render();
                // 选择 svg 并添加一个g元素作为绘图容器.
                let svg = d3.select("#svg-canvas"); //声明节点
                svg.select("g").remove(); //删除以前的节点，清空画面
                let svgGroup = svg.append("g");
                let inner = svg.select("g");
                // 在绘图容器上运行渲染器生成流程图.
                let zoom = d3.zoom().on("zoom", function() { //添加鼠标滚轮放大缩小事件
                    inner.attr("transform", d3.event.transform);
                });
                svg.call(zoom);
                render(svgGroup, g);

                d3.selectAll("g.node").on("click", function(id) {
                    $that._loadReqInfo(id);
                });

                let max = svg._groups[0][0].clientWidth > svg._groups[0][0].clientHeight ? svg._groups[0][0].clientWidth : svg._groups[0][0].clientHeight;
                let initialScale = max / 779; //initialScale元素放大倍数，随着父元素宽高发生变化时改变初始渲染大小
                let tWidth = (svg._groups[0][0].clientWidth - g.graph().width * initialScale) / 2; //水平居中
                let tHeight = (svg._groups[0][0].clientHeight - g.graph().height * initialScale) / 2; //垂直居中
                svgGroup.call(zoom.transform, d3.zoomIdentity.translate(tWidth, tHeight).scale(initialScale)); //元素水平垂直居中复制代码
            },
            _goBack: function() {
                vc.goBack();
            },
            _loadReqInfo: function(id) {
                let _traces = $that.logTraceDetailInfo.traces;
                let _curTrace = {};
                _traces.forEach((item) => {
                    if (item.id == id) {
                        _curTrace = item;
                    }
                });
                vc.copyObject(_curTrace, $that.logTraceDetailInfo)
                $that.logTraceDetailInfo.annos = _curTrace.annotations;
                $that._listLogTraceParam();
            },
            _listLogTraceParam: function() {
                let param = {
                    params: {
                        spanId: $that.logTraceDetailInfo.spanId,
                        page: 1,
                        row: 1
                    }
                };
                //发送get请求
                vc.http.apiGet('/monitor/getLogTraceParam',
                    param,
                    function(json, res) {
                        let _logTraceInfo = JSON.parse(json);

                        if (_logTraceInfo.data.length < 1) {
                            return;
                        }
                        vc.copyObject(logTraceInfo.data[0], $that.logTraceDetailInfo)
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },

            _getAnnoName: function(_value) {
                if (_value == "cs") {
                    return '接受时间'
                } else if (_value == "ss") {
                    return '调用下游';
                } else if (_value == "sr") {
                    return '接受下游';
                } else {
                    return '返回时间';
                }

            },
            _getAnnoTime: function(_time) {
                return vc.dateTimeFormat(_time)
            },
            compare:function (obj1, obj2) {
                let val1 = obj1.timestamp;
                let val2 = obj2.timestamp;
                if (val1 < val2) {
                    return -1;
                } else if (val1 > val2) {
                    return 1;
                } else {
                    return 0;
                }            
            } 

        }
    });
})(window.vc);