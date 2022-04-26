
getParam = function (_key) {
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

initWindow = function () {
    let _zihaoToken = getParam('zihaoToken');
    let _hostId = getParam('hostId');
    let winWidth = document.documentElement.clientWidth;
    let winHeight = document.documentElement.clientHeight;
    let protocol = "";
    if (window.location.protocol == 'https:') {
         protocol = 'wss://';
    } else {
         protocol = 'ws://';
    }
    let endpoint = protocol + window.location.host + '/app/console/webWindow';
    let display = document.getElementById("display");
    // Instantiate client, using an HTTP tunnel for communications.
    let guac = new Guacamole.Client(
        new Guacamole.WebSocketTunnel(endpoint)
    );
    // Add client to display div
    display.appendChild(guac.getDisplay().getElement());
    // Error handler
    guac.onerror = function (error) {
        alert(error);
    };
    // Connect
    let _q = 'zihaoToken='
    +_zihaoToken+"&hostId="+_hostId+"&winWidth="+winWidth+"&winHeight="+winHeight
    guac.connect(_q);

    // Disconnect on close
    window.onunload = function () {
        guac.disconnect();
    }

    // Mouse
    let mouse = new Guacamole.Mouse(guac.getDisplay().getElement());

    mouse.onEach(['mousedown', 'mouseup', 'mousemove'], function sendMouseEvent(e) {
        guac.sendMouseState(e.state);
    });

    // Keyboard
    let keyboard = new Guacamole.Keyboard(document);

    keyboard.onkeydown = function (keysym) {
        guac.sendKeyEvent(1, keysym);
    };

    keyboard.onkeyup = function (keysym) {
        guac.sendKeyEvent(0, keysym);
    };
}

initWindow();
   
