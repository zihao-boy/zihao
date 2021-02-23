/**
 导航栏
 **/
(function (vc) {
    let DEFAULT_PAGE = 1;
    let DEFAULT_ROW = 10;
    var vm = new Vue({
        el: '#nav',
        data: {
            nav: {
                moreNoticeUrl: '/admin.html#/pages/admin/noticeManage',
                notices: [],
                total: 0,
                _currentCommunity: '',
                communityInfos: []
            },
            logo: '',
            userName: "",

        },
        mounted: function () {
            this._initSysInfo();
            //this.getNavData();
            //this.getUserInfo();
        },
        methods: {
            _initSysInfo: function () {
                var sysInfo = vc.getData("_sysInfo");
                if (sysInfo == null) {
                    this.logo = "HC";
                    return;
                }
                this.logo = sysInfo.logo;
            },
            getNavData: function () {

                var param = {
                    params: {
                        page: 1,
                        row: 3
                    }

                };

                //发送get请求
                vc.http.get('nav',
                    'getNavData',
                    param,
                    function (json) {
                        var _noticeObj = JSON.parse(json);
                        vm.nav.notices = _noticeObj.msgs;
                        vm.nav.total = _noticeObj.total;
                    }, function () {
                        console.log('请求失败处理');
                    }
                );

            },
            logout: function () {
                var param = {
                    msg: 123
                };
                //发送get请求
                vc.http.apiPost('/user/logout',
                    JSON.stringify(param),
                    {
                        emulateJSON: true
                    },
                    function (json, res) {
                        if (res.status == 200) {
                            vc.jumpToPage("/user.html#/pages/frame/login");
                            return;
                        }
                    }, function () {
                        console.log('请求失败处理');
                    }
                );
            },
            getUserInfo: function () {
                //获取用户名
                var param = {
                    msg: '123',
                };

                //发送get请求
                vc.http.apiGet('/user/getUserInfo',
                    param,
                    function (json, res) {
                    
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            vm.userName = _json.data.realName;
                            vc.saveData('/nav/getUserInfo',_json.data)
                            return;
                        }
                        vc.toast(_json.msg);
                    }, function () {
                        console.log('请求失败处理');
                    }
                );
            },
            
        
            _noticeDetail: function (_msg) {
                //console.log(_notice.noticeId);
                //vc.jumpToPage("/admin.html#/noticeDetail?noticeId="+_notice.noticeId);

                //标记为消息已读
                vc.http.post('nav',
                    'readMsg',
                    JSON.stringify(_msg),
                    function (json, res) {
                        if (res.status == 200) {
                            vc.jumpToPage(_msg.url);
                        }
                    }, function () {
                        console.log('请求失败处理');
                    }
                );
            },
            _doMenu: function () {
                let body = document.getElementsByTagName("body")[0];

                let className = body.className;

                if (className.indexOf("mini-navbar") != -1) {
                    body.className = className.replace(/mini-navbar/g, "");
                    return;
                }
                body.className = className + " mini-navbar";
            },
            _chooseMoreCommunity: function () {
                vc.emit('chooseEnterCommunity', 'openChooseEnterCommunityModel', {});
            },
            _viewDocument:function(){
                vc.emit('document','openDocument', {});
            }
        }
    });

    vm.getUserInfo();

    //建立websocket 消息连接
    let _userId = vc.getData('/nav/getUserInfo').userId;

    let _protocol = window.location.protocol;
    let url = '';
    if (_protocol.startsWith('https')) {
        url =
            "wss://" + window.location.host + "/app/message/" +
            _userId;
    } else {
        url =
            "ws://" + window.location.host + "/app/message/" +
            _userId;
    }


    if ("WebSocket" in window) {
        websocket = new WebSocket(url);
    } else if ("MozWebSocket" in window) {
        websocket = new MozWebSocket(url);
    } else {
        websocket = new SockJS(url);
    }

    //连接发生错误的回调方法
    websocket.onerror = function () {
        console.log("初始化失败");
        this.$notify.error({
            title: "错误",
            message: "连接失败，请检查网络"
        });
    };

    //连接成功建立的回调方法
    websocket.onopen = function () {
        console.log("ws初始化成功");
    };

    //接收到消息的回调方法
    websocket.onmessage = function (event) {
        console.log("event", event);
        //let _data = event.data;
        let _data = JSON.parse(_data);
        if (_data.code == 200) {
            toastr.info(_data.msg);
        } else {
            toastr.error(_data.msg);
        }
    };

    //连接关闭的回调方法
    websocket.onclose = function () {
        console.log("初始化失败");
        this.$notify.error({
            title: "错误",
            message: "连接关闭，请刷新浏览器"
        });
    };

    //监听窗口关闭事件，当窗口关闭时，主动去关闭websocket连接，防止连接还没断开就关闭窗口，server端会抛异常。
    window.onbeforeunload = function () {
        websocket.close();
    };

})(window.vc);