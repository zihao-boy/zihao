$(function() {

    $(window).resize(function() {
        hightChange();
    });

    function hightChange() {　　
        var h = document.documentElement.clientHeight - 160;　　
        $("#echartsMap").height(h); // iframe id
        $(".data_main").height(h); // iframe id
    }

    hightChange();
    getAccessLogMap();
    getAccessLogTop5();
    getAccessLogIntercept();

    setInterval(function() {
        getAccessLogMap();
        getAccessLogTop5();
        getAccessLogIntercept();
    }, 10000)


    document.getElementById('echartsMap').style.height = clientHeight

    function getAccessLogMap() {
        let xhr = new XMLHttpRequest();
        let _QData = [];
        let _CData = "北京";

        //第二步 打开要发送的地址
        xhr.open("get", "/app/firewall/getWafAccessLogMap");
        //第三部发送
        xhr.send();
        //第四步
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 && xhr.status == 200) {
                let _data = JSON.parse(xhr.responseText);
                if (!_data.data) {
                    echart_map(_CData, _QData);
                    return;
                }

                let _location = getLocation();
                for (let key in _location) {
                    _data.data.forEach(_item => {
                        if (_item.country.indexOf(key) > -1) {
                            _item.name = key;
                            _QData.push(_item);
                        }
                        if (_item.wafCountry.indexOf(key) > -1) {
                            _CData = key;
                        }
                    })
                }
                echart_map(_CData, _QData);
            }
        }
    }

    function getAccessLogTop5() {
        let xhr = new XMLHttpRequest();
        let _QData = [];
        let _CData = "北京";

        //第二步 打开要发送的地址
        xhr.open("get", "/app/firewall/getWafAccessLogTop5");
        //第三部发送
        xhr.send();
        //第四步
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 && xhr.status == 200) {
                let _data = JSON.parse(xhr.responseText);
                if (!_data.data) {
                    return;
                }
                let _html = "";
                _data.data.forEach(_item => {
                    _html += ("<tr>" +
                        "<td>" + _item.country + "</td>" +
                        "<td>" + _item.xRealIp + "</td>" +
                        "<td>" + _item.total + "</td>" +
                        "</tr>"
                    )
                });
                document.getElementById('accessLogTop5').innerHTML = _html;
            }
        }
    }

    function getAccessLogIntercept() {
        let xhr = new XMLHttpRequest();

        //第二步 打开要发送的地址
        xhr.open("get", "/app/firewall/getWafAccessLogIntercept");
        //第三部发送
        xhr.send();
        //第四步
        xhr.onreadystatechange = function() {
            if (xhr.readyState == 4 && xhr.status == 200) {
                let _data = JSON.parse(xhr.responseText);
                if (!_data.data) {
                    return;
                }
                let _html = "";
                _data.data.forEach(_item => {
                    _html += ("<tr>" +
                        "<td>" + _item.country + "</td>" +
                        "<td>" + _item.xRealIp + "</td>" +
                        "<td>" + _item.stateName + "</td>" +
                        "<td>" + formateTime(_item.createTime) + "</td>" +
                        "</tr>"
                    )
                });
                document.getElementById('intercept').innerHTML = _html;
            }
        }
    }

    function formateTime(_dateStr) {
        let _date = _dateStr.replace('T', ' ').replace('Z', ' ');
        let newDate = new Date(_date);
        let m = newDate.getMonth() + 1;
        let d = newDate.getDate();
        let h = newDate.getHours();
        let mm = newDate.getMinutes();
        return m + '-' + d + ' ' + h + ':' + mm;
    }

    // echart_map中国地图
    function echart_map(_CData, _QData) {
        let chart = echarts.init(document.getElementById('echartsMap'));
        /*
            图中相关城市经纬度,根据你的需求添加数据
            关于国家的经纬度，可以用首都的经纬度或者其他城市的经纬度
        */
        var geoCoordMap = getLocation();
        /* 
            记录下起点城市和终点城市的名称，以及权重
            数组第一位为终点城市，数组第二位为起点城市，以及该城市的权重，就是图上圆圈的大小
         */
        // 重庆
        let CQData = [];
        let _ICQData = [];
        _QData.forEach(_item => {
            _ICQData = [];
            _ICQData.push({
                name: _CData
            });
            _ICQData.push({
                name: _item.name,
                value: 30
            });
            CQData.push(_ICQData)
        });



        // 小飞机的图标，可以用其他图形替换
        //var planePath = 'path://M1705.06,1318.313v-89.254l-319.9-221.799l0.073-208.063c0.521-84.662-26.629-121.796-63.961-121.491c-37.332-0.305-64.482,36.829-63.961,121.491l0.073,208.063l-319.9,221.799v89.254l330.343-157.288l12.238,241.308l-134.449,92.931l0.531,42.034l175.125-42.917l175.125,42.917l0.531-42.034l-134.449-92.931l12.238-241.308L1705.06,1318.313z';
        var planePath = 'image://waf/img/echart_come.png'
            // 获取地图中起点和终点的坐标，以数组形式保存下来
        var convertData = function(data) {
            var res = [];
            for (var i = 0; i < data.length; i++) {
                var dataItem = data[i];
                var fromCoord = geoCoordMap[dataItem[1].name];
                var toCoord = geoCoordMap[dataItem[0].name];
                if (fromCoord && toCoord) {
                    res.push([{
                        coord: fromCoord // 起点坐标
                    }, {
                        coord: toCoord // 终点坐标
                    }])
                }
            }
            return res;
        }

        var color = ['#9ae5fc', '#dcbf71']; // 自定义图中要用到的颜色
        var series = []; // 用来存储地图数据

        /*
            图中一共用到三种效果，分别为航线特效图、飞机航线图以及城市图标涟漪图。
            要用到setOption中的series属性，并且对每个城市都要进行三次设置。
        */
        [
            [_CData, CQData]
        ].forEach(function(item, i) {
            series.push({
                // 白色航线特效图
                type: 'lines',
                zlevel: 1, // 用于分层，z-index的效果
                effect: {
                    show: true, // 动效是否显示
                    period: 6, // 特效动画时间
                    trailLength: 0.7, // 特效尾迹的长度
                    color: '#fff', // 特效颜色
                    symbolSize: 3 // 特效大小
                },
                lineStyle: {
                    normal: { // 正常情况下的线条样式
                        color: color[0],
                        width: 0, // 因为是叠加效果，要是有宽度，线条会变粗，白色航线特效不明显
                        curveness: -0.2 // 线条曲度
                    }
                },
                data: convertData(item[1]) // 特效的起始、终点位置
            }, { // 小飞机航线效果
                type: 'lines',
                zlevel: 2,
                //symbol: ['none', 'arrow'],   // 用于设置箭头
                symbolSize: 10,
                effect: {
                    show: true,
                    period: 6,
                    trailLength: 0,
                    symbol: planePath, // 特效形状，可以用其他svg pathdata路径代替
                    symbolSize: 15
                },
                lineStyle: {
                    normal: {
                        color: color[0],
                        width: 1,
                        opacity: 0.6,
                        curveness: -0.2
                    }
                },
                data: convertData(item[1]) // 特效的起始、终点位置，一个二维数组，相当于coords: convertData(item[1])
            }, { // 散点效果
                type: 'effectScatter',
                coordinateSystem: 'geo', // 表示使用的坐标系为地理坐标系
                zlevel: 3,
                rippleEffect: {
                    brushType: 'stroke' // 波纹绘制效果
                },
                label: {
                    normal: { // 默认的文本标签显示样式
                        show: true,
                        position: 'left', // 标签显示的位置
                        formatter: '{b}' // 标签内容格式器
                    }
                },
                itemStyle: {
                    normal: {
                        color: color[0]
                    }
                },
                data: item[1].map(function(dataItem) {
                    return {
                        name: dataItem[1].name,
                        value: geoCoordMap[dataItem[1].name], // 起点的位置
                        symbolSize: dataItem[1].value / 8, // 散点的大小，通过之前设置的权重来计算，val的值来自data返回的value
                    };
                })
            });
        });

        // 显示终点位置,类似于上面最后一个效果，放在外面写，是为了防止被循环执行多次
        series.push({
            type: 'effectScatter',
            coordinateSystem: 'geo',
            zlevel: 3,
            rippleEffect: {
                brushType: 'stroke'
            },
            label: {
                normal: {
                    show: true,
                    position: 'left',
                    formatter: '{b}'
                }
            },
            symbolSize: function(val) {
                return val[2] / 8;
            },
            itemStyle: {
                normal: {
                    color: color[1]
                }
            },
            data: [{
                // 这里面的数据，由于一开始就知道终点位置是什么，所以直接写死，如果通过ajax来获取数据的话，还要进行相应的处理
                name: _CData,
                value: geoCoordMap[_CData],
                label: {
                    normal: {
                        position: 'top'
                    }
                }
            }]
        });

        // 最后初始化世界地图中的相关数据
        chart.setOption({
            title: {
                text: '',
                textStyle: {
                    color: '#fff',
                    fontSize: 40
                },
                top: '10px',
                left: '10px'
            },
            geo: {
                map: 'world', // 与引用进来的地图js名字一致
                roam: true, // 禁止缩放平移
                center: [108.95000, 34.26667],
                zoom: 5,
                itemStyle: { // 每个区域的样式 
                    normal: {
                        areaColor: '#323c48'
                    },
                    emphasis: {
                        areaColor: '#2a333d'
                    }
                },
                regions: [{ // 选中的区域
                    name: 'China',
                    selected: true,
                    itemStyle: { // 高亮时候的样式
                        emphasis: {
                            areaColor: '#7d7d7d'
                        }
                    },
                    label: { // 高亮的时候不显示标签
                        emphasis: {
                            show: false
                        }
                    }
                }]
            },
            series: series, // 将之前处理的数据放到这里
            textStyle: {
                fontSize: 12
            }
        });
        window.addEventListener("resize", function() {
            chart.resize();
        });
    }



    function getLocation() {
        return {
            '山东': [117.000923, 36.675807],
            '河北': [115.48333, 38.03333],
            '吉林': [125.35000, 43.88333],
            '黑龙江': [127.63333, 47.75000],
            '辽宁': [123.38333, 41.80000],
            '内蒙古': [111.670801, 41.818311],
            '新疆': [87.68333, 43.76667],
            '甘肃': [103.73333, 36.03333],
            '宁夏': [106.26667, 37.46667],
            '山西': [112.53333, 37.86667],
            '陕西': [108.95000, 34.26667],
            '河南': [113.65000, 34.76667],
            '安徽': [117.283042, 31.86119],
            '江苏': [119.78333, 32.05000],
            '浙江': [120.20000, 30.26667],
            '福建': [118.30000, 26.08333],
            '广东': [113.23333, 23.16667],
            '江西': [115.90000, 28.68333],
            '海南': [110.35000, 20.01667],
            '广西': [108.320004, 22.82402],
            '贵州': [106.71667, 26.56667],
            '湖南': [113.00000, 28.21667],
            '湖北': [114.298572, 30.584355],
            '四川': [104.06667, 30.66667],
            '云南': [102.73333, 25.05000],
            '西藏': [91.00000, 30.60000],
            '青海': [96.75000, 36.56667],
            '天津': [117.20000, 39.13333],
            '上海': [121.55333, 31.20000],
            '重庆': [106.45000, 29.56667],
            '北京': [116.41667, 39.91667],
            '台湾': [121.30, 25.03],
            '香港': [114.10000, 22.20000],
            '澳门': [113.50000, 22.20000],
            "阿富汗": [69.11, 34.28],
            "阿尔巴尼亚": [19.49, 41.18],
            "阿尔及利亚": [3.08, 36.42],
            "美属萨摩亚": [-170.43, -14.16],
            "安道​​尔": [1.32, 42.31],
            "安哥拉": [13.15, -8.50],
            "安提瓜和巴布达": [-61.48, 17.20],
            "阿根廷": [-60.00, -36.30],
            "亚美尼亚": [44.31, 40.10],
            "阿鲁巴": [-70.02, 12.32],
            "澳大利亚": [149.08, -35.15],
            "奥地利": [16.22, 48.12],
            "阿塞拜疆": [49.56, 40.29],
            "巴哈马": [-77.20, 25.05],
            "巴林": [50.30, 26.10],
            "孟加拉国": [90.26, 23.43],
            "巴巴多斯": [-59.30, 13.05],
            "白俄罗斯": [27.30, 53.52],
            "比利时": [4.21, 50.51],
            "伯利兹": [-88.30, 17.18],

            "贝宁": [2.42, 6.23],

            "不丹": [89.45, 27.31],

            "玻利维亚": [-68.10, -16.20],

            "波斯尼亚和黑塞哥维那": [18.26, 43.52],

            "博茨瓦纳": [25.57, -24.45],

            "巴西": [-47.55, -15.47],

            "英属维尔京群岛": [-64.37, 18.27],

            "文莱": [115.00, 4.52],

            "保加利亚": [23.20, 42.45],

            "布基纳法索": [-1.30, 12.15],

            "布隆迪": [29.18, -3.16],

            "柬埔寨": [104.55, 11.33],

            "喀麦隆": [11.35, 3.50],

            "加拿大": [-75.42, 45.27],

            "佛得角": [-23.34, 15.02],

            "开曼群岛": [-81.24, 19.20],

            "中非共和国": [18.35, 4.23],

            "乍得": [14.59, 12.10],

            "智利": [-70.40, -33.24],

            "中国": [116.20, 39.55],

            "哥伦比亚": [-74.00, 4.34],

            "科摩罗": [43.16, -11.40],

            "刚果": [15.12, -4.09],

            "哥斯达黎加": [-84.02, 9.55],

            "科特迪瓦": [-5.17, 6.49],

            "克罗地亚": [15.58, 45.50],

            "古巴": [-82.22, 23.08],

            "塞浦路斯": [33.25, 35.10],

            "捷克共和国": [14.22, 50.05],

            "朝鲜": [125.30, 39.09],

            "刚果(扎伊尔)": [15.15, -4.20],

            "丹麦": [12.34, 55.41],

            "吉布提": [42.20, 11.08],

            "多米尼加": [-61.24, 15.20],

            "多米尼加共和国": [-69.59, 18.30],

            "东帝汶": [125.34, -8.29],

            "厄瓜多尔": [-78.35, -0.15],

            "埃及": [31.14, 30.01],

            "萨尔瓦多": [-89.10, 13.40],
            "赤道几内亚": [8.50, 3.45],
            "厄立特里亚": [38.55, 15.19],
            "爱沙尼亚": [24.48, 59.22],
            "埃塞俄比亚": [38.42, 9.02],
            "福克兰群岛(马尔维纳斯群岛)": [-59.51, -51.40],
            "法罗群岛": [-6.56, 62.05],
            "斐济": [178.30, -18.06],

            "芬兰": [25.03, 60.15],

            "法国": [2.20, 48.50],

            "法属圭亚那": [-52.18, 5.05],

            "法属波利尼西亚": [-149.34, -17.32],

            "加蓬": [9.26, 0.25],

            "冈比亚": [-16.40, 13.28],

            "格鲁吉亚": [44.50, 41.43],

            "德国": [13.25, 52.30],

            "加纳": [-0.06, 5.35],

            "希腊": [23.46, 37.58],

            "格陵兰": [-51.35, 64.10],

            "瓜德罗普岛": [-61.44, 16.00],

            "危地马拉": [-90.22, 14.40],

            "根西岛": [-2.33, 49.26],

            "几内亚": [-13.49, 9.29],

            "几内亚比绍": [-15.45, 11.45],

            "圭亚那": [-58.12, 6.50],

            "海地": [-72.20, 18.40],

            "赫德岛和麦当劳群岛": [74.00, -53.00],

            "洪都拉斯": [-87.14, 14.05],

            "匈牙利": [19.05, 47.29],

            "冰岛": [-21.57, 64.10],

            "印度": [77.13, 28.37],

            "印度尼西亚": [106.49, -6.09],

            "伊朗": [51.30, 35.44],

            "伊拉克": [44.30, 33.20],

            "爱尔兰": [-6.15, 53.21],

            "以色列": [35.12, 31.47],

            "意大利": [12.29, 41.54],

            "牙买加": [-76.50, 18.00],

            "约旦": [35.52, 31.57],

            "哈萨克斯坦": [71.30, 51.10],

            "肯尼亚": [36.48, -1.17],

            "基里巴斯": [173.00, 1.30],

            "科威特": [48.00, 29.30],

            "吉尔吉斯斯坦": [74.46, 42.54],

            "老挝": [102.36, 17.58],

            "拉脱维亚": [24.08, 56.53],

            "黎巴嫩": [35.31, 33.53],

            "莱索托": [27.30, -29.18],

            "利比里亚": [-10.47, 6.18],

            "阿拉伯利比亚民众国": [13.07, 32.49],

            "列支敦士登": [9.31, 47.08],

            "立陶宛": [25.19, 54.38],

            "卢森堡": [6.09, 49.37],

            "马达加斯加": [47.31, -18.55],

            "马拉维": [33.48, -14.00],

            "马来西亚": [101.41, 3.09],

            "马尔代夫": [73.28, 4.00],

            "马里": [-7.55, 12.34],

            "马耳他": [14.31, 35.54],

            "马提尼克岛": [-61.02, 14.36],

            "毛里塔尼亚": [57.30, -20.10],

            "马约特岛": [45.14, -12.48],

            "墨西哥": [-99.10, 19.20],

            "密克罗尼西亚(联邦) ": [158.09, 6.55],

            "摩尔多瓦共和国": [28.50, 47.02],

            "莫桑比克": [32.32, -25.58],

            "缅甸": [96.20, 16.45],

            "纳米比亚": [17.04, -22.35],

            "尼泊尔": [85.20, 27.45],
            "荷兰": [4.54, 52.23],

            "荷属安的列斯": [-69.00, 12.05],

            "新喀里多尼亚": [166.30, -22.17],

            "新西兰": [174.46, -41.19],

            "尼加拉瓜": [-86.20, 12.06],

            "尼日尔": [2.06, 13.27],

            "尼日利亚": [7.32, 9.05],

            "诺福克岛": [168.43, -45.20],

            "北马里亚纳群岛": [145.45, 15.12],

            "挪威": [10.45, 59.55],

            "阿曼": [58.36, 23.37],

            "巴基斯坦": [73.10, 33.40],

            "帕劳": [134.28, 7.20],

            "巴拿马": [-79.25, 9.00],

            "巴布亚新几内亚": [147.08, -9.24],

            "巴拉圭": [-57.30, -25.10],

            "秘鲁": [-77.00, -12.00],

            "菲律宾": [121.03, 14.40],

            "波兰": [21.00, 52.13],

            "葡萄牙": [-9.10, 38.42],

            "波多黎各": [-66.07, 18.28],

            "卡塔尔": [51.35, 25.15],

            "韩国": [126.58, 37.31],

            "罗马尼亚": [26.10, 44.27],

            "俄罗斯": [37.35, 55.45],

            "卢旺达": [30.04, -1.59],

            "圣基茨和尼维斯": [-62.43, 17.17],

            "圣卢西亚": [-60.58, 14.02],

            "圣皮埃尔和密克隆": [-56.12, 46.46],

            "圣文森特和格林纳丁斯": [-61.10, 13.10],

            "萨摩亚": [-171.50, -13.50],

            "圣马力诺": [12.30, 43.55],

            "圣多美和普林西比": [6.39, 0.10],

            "沙特阿拉伯": [46.42, 24.41],

            "塞内加尔": [-17.29, 14.34],

            "塞拉利昂": [-13.17, 8.30],

            "斯洛伐克": [17.07, 48.10],

            "斯洛文尼亚": [14.33, 46.04],

            "所罗门群岛": [159.57, -9.27],

            "索马里": [45.25, 2.02],

            "比勒陀利亚": [28.12, -25.44],

            "西班牙": [-3.45, 40.25],

            "苏丹": [32.35, 15.31],

            "苏里南": [-55.10, 5.50],

            "斯威士兰": [31.06, -26.18],

            "瑞典": [18.03, 59.20],

            "瑞士": [7.28, 46.57],

            "阿拉伯叙利亚共和国": [36.18, 33.30],

            "塔吉克斯坦": [68.48, 38.33],

            "泰国": [100.35, 13.45],

            "马其顿": [21.26, 42.01],

            "多哥": [1.20, 6.09],

            "汤加": [-174.00, -21.10],

            "突尼斯": [10.11, 36.50],

            "土耳其": [32.54, 39.57],

            "土库曼斯坦": [57.50, 38.00],

            "图瓦卢": [179.13, -8.31],

            "乌干达": [32.30, 0.20],

            "乌克兰": [30.28, 50.30],

            "阿联酋": [54.22, 24.28],

            "英国": [-0.05, 51.36],

            "坦桑尼亚": [35.45, -6.08],

            "美国": [-77.02, 39.91],

            "美属维尔京群岛": [-64.56, 18.21],

            "乌拉圭": [-56.11, -34.50],

            "乌兹别克斯坦": [69.10, 41.20],

            "瓦努阿图": [168.18, -17.45],

            "委内瑞拉": [-66.55, 10.30],

            "越南": [105.55, 21.05],

            "南斯拉夫": [20.37, 44.50],
            "赞比亚": [28.16, -15.28],
            "津巴布韦": [31.02, -17.43]
        };
    }
});