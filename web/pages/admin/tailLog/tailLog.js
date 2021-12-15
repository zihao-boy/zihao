/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            tailLogInfo: {
                log: '',
                path:''
                
            }
        },
        _initMethod: function () {
            $that.tailLogInfo.path = vc.getParam('path');
            vc.component._initWs();
        },
        _initEvent: function () {

        },
        methods: {
            _initWs: function (_page, _rows) {
                let _connection = null;
                let protocol = 'ws://';
                if (window.location.protocol == 'https:') {
                     protocol = 'wss://';
                } 
                //window.location.host
                let endpoint = protocol + window.location.host + '/app/console/tailLog';
                if (window.WebSocket) {
                    //如果支持websocket
                    _connection = new WebSocket(endpoint);
                } else {
                    //否则报错
                    options.onError('WebSocket Not Supported');
                    return;
                }

                _connection.onopen = function () {
                    _connection.send($that.tailLogInfo.path);
                };
            
                _connection.onmessage = function (evt) {
                    let data = evt.data.toString();
                    $that.tailLogInfo.log += data;
                };
            
            
                _connection.onclose = function (evt) {
                    vc.toast('日志关闭,请刷新重新打开！');
                };

            },
            _goBack:function(){
                vc.goBack();
            }
        }
    });
})(window.vc);