package defaultWebsocket

import (
	"github.com/zihao-boy/zihao/common/webLog"
	"github.com/zihao-boy/zihao/common/webWindow"
	"log"
	"net/http"
	"strconv"

	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"github.com/zihao-boy/zihao/common/webshell"
)

func InitWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}), websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())

			webshell.WebSocketHandler(msg.Body, nsConn.Conn.ID(), nsConn)

			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {

		log.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		log.Printf("[%s] Disconnected from server", c.ID())
		webshell.CloseSsh(c.ID())
	}

	ws.OnUpgradeError = func(err error) {
		log.Printf("Upgrade Error: %v", err)
	}

	app.Get("/app/console/ssh", websocket.Handler(ws))
}


func InitWebsocketWindow(app *iris.Application) {
	// create our echo websocket server
	var reqData webWindow.Request
	upgrader :=gorillaWs.Upgrader{CheckOrigin: func(req *http.Request) bool {
		values := req.URL.Query()
		winWidth,_:=strconv.Atoi(values.Get("winWidth"))
		winHeight,_:=strconv.Atoi(values.Get("winHeight"))
		reqData = webWindow.Request{
			ZihaoToken:values.Get("zihaoToken"),
			HostId:values.Get("hostId"),
			WinWidth:winWidth,
			WinHeight:winHeight,
		}
		return true
	},Subprotocols: []string{"guacamole"}}

	ws := websocket.New(gorilla.Upgrader(upgrader), websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
			webWindow.WebSocketHandler(msg.Body, nsConn.Conn.ID(), nsConn)
			return nil
		},
	})




	ws.OnConnect = func(c *websocket.Conn) error {
		log.Printf("[%s] Connected to server!", c.ID())
		webWindow.WebSocketConn(reqData,c.ID(),c)
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		log.Printf("[%s] Disconnected from server", c.ID())
		webWindow.CloseSsh(c.ID())
	}

	ws.OnUpgradeError = func(err error) {
		log.Printf("Upgrade Error: %v", err)
	}

	app.Get("/app/console/webWindow", websocket.Handler(ws))
}


func InitLogWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool { return true }}), websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())

			webLog.WebSocketHandler(msg.Body, nsConn.Conn.ID(), nsConn)

			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {

		log.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		log.Printf("[%s] Disconnected from server", c.ID())
		webLog.CloseTail(c.ID())
	}

	ws.OnUpgradeError = func(err error) {
		log.Printf("Upgrade Error: %v", err)
	}

	app.Get("/app/console/tailLog", websocket.Handler(ws))
}


