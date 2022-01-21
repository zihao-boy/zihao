function WSSHClient() {};

WSSHClient.prototype._generateEndpoint = function() {
    if (window.location.protocol == 'https:') {
        var protocol = 'wss://';
    } else {
        var protocol = 'ws://';
    }
    //window.location.host
    //var endpoint = protocol+'localhost:7000/app/console/ssh';
    //var endpoint = protocol+'192.168.1.34:7000/app/console/ssh';
    var endpoint = protocol + window.location.host + '/app/console/ssh';
    return endpoint;
};

WSSHClient.prototype.connect = function(options) {
    var endpoint = this._generateEndpoint();

    if (window.WebSocket) {
        //如果支持websocket
        this._connection = new WebSocket(endpoint);
    } else {
        //否则报错
        options.onError('WebSocket Not Supported');
        return;
    }

    this._connection.onopen = function() {
        options.onConnect();
    };

    this._connection.onmessage = function(evt) {
        var data = evt.data.toString();
        //data = base64.decode(data);
        options.onData(data);
    };


    this._connection.onclose = function(evt) {
        options.onClose();
    };
};

WSSHClient.prototype.send = function(data) {
    this._connection.send(JSON.stringify(data));
};

WSSHClient.prototype.sendInitData = function(options) {
    //连接参数
    this._connection.send(JSON.stringify(options));
}

WSSHClient.prototype.sendClientData = function(data) {
    //发送指令
    this._connection.send(JSON.stringify({ "operate": "command", "command": data }))
}

var client = new WSSHClient();

getParam = function(_key) {
    //返回当前 URL 的查询部分（问号 ? 之后的部分）。
    let urlParameters = location.search;
    if (urlParameters != null && urlParameters != undefined && urlParameters != '') {
        //urlParameters = location.hash;

        if (urlParameters.indexOf('?') != -1) {
            urlParameters = urlParameters.substring(urlParameters.indexOf('?'), urlParameters.length);
        }
    }
    //如果该求青中有请求的参数，则获取请求的参数，否则打印提示此请求没有请求的参数
    if (urlParameters.indexOf('?') != -1) {
        //获取请求参数的字符串
        let parameters = decodeURI(urlParameters.substr(1));
        //将请求的参数以&分割中字符串数组
        parameterArray = parameters.split('&');
        //循环遍历，将请求的参数封装到请求参数的对象之中
        for (let i = 0; i < parameterArray.length; i++) {
            if (_key == parameterArray[i].split('=')[0]) {
                return parameterArray[i].split('=')[1];
            }
        }
    }
    return "";
};


function initShell() {

    let _zihaoToken = getParam('zihaoToken');
    let _hostId = getParam('hostId');
    let command = getParam('command');
    let val = getParam('val');

    if (command != null && val == 'log') {
        command = 'docker logs -f -t --tail 1 ' + command;
    }
    if (command != null && val == 'exec') {

        command = 'docker exec -it ' + command + ' /bin/bash';
    }

    if (command != null && val == 'restart') {

        command = 'docker restart ' + command + '\n docker logs -f --tail 1 ' + command;
    }

    if (command != null && val == 'cd') {

        command = 'cd ' + command;
    }
    //document.body.clientWidth ||
    let winWidth = document.documentElement.clientWidth;
    //document.body.clientHeight ||
    let winHeight = document.documentElement.clientHeight;
    let _cols = Math.floor(winWidth / 9) - 5;
    let _rows = Math.floor(winHeight / 16) - 4;

    openTerminal({
        operate: 'connect',
        zihaoToken: _zihaoToken,
        command: command,
        hostId: _hostId,
        winWidth: winWidth,
        winHeight: winHeight,
        cols: _cols,
        rows: _rows,
    });

    function openTerminal(options) {
        var client = new WSSHClient();

        var term = new Terminal({
            cols: _cols,
            rows: _rows,
            cursorBlink: true, // 光标闪烁
            cursorStyle: "block", // 光标样式  null | 'block' | 'underline' | 'bar'
            scrollback: 800, //回滚
            tabStopWidth: 30, //制表宽度
            screenKeys: true,
            winHeight: winHeight,
            winWidth: winWidth
        });

        term.on('data', function(data) {
            //键盘输入时的回调函数
            client.sendClientData(data);
        });
        term.open(document.getElementById('terminal'));
        //在页面上显示连接中...
        term.write('connection...');
        //执行连接操作
        client.connect({
            onError: function(error) {
                //连接失败回调
                term.write('error: ' + error + '\r\n');
            },
            onConnect: function() {
                //连接成功回调
                client.sendInitData(options);
                if (command != null && command != undefined && command != '') {
                    client.sendClientData(command + '\n')
                }
            },
            onClose: function() {
                //连接关闭回调
                term.write("\connection close ....");
            },
            onData: function(data) {
                //收到数据时回调
                term.write(data);
            }
        });
    }

}

initShell()