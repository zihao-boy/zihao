package defaultWebsocket

import (
	gorillaWs "github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gorilla"
	"github.com/zihao-boy/zihao/zihao-service/common/webshell"
	"log"
	"net/http"
)

func InitWebsocket(app *iris.Application){
	// create our echo websocket server
	ws := websocket.New(gorilla.Upgrader(gorillaWs.Upgrader{CheckOrigin: func(*http.Request) bool{return true}}), websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())

			webshell.WebSocketHandler(msg.Body, func(outParam []byte) {
				mg := websocket.Message{
					Body:outParam,
					IsNative:true,
				}

				nsConn.Conn.Write(mg)
			})


			return nil
		},
	})


	ws.OnConnect = func(c *websocket.Conn) error {
		log.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		log.Printf("[%s] Disconnected from server", c.ID())
	}

	ws.OnUpgradeError = func(err error) {
		log.Printf("Upgrade Error: %v", err)
	}

	app.Get("/app/console/ssh",websocket.Handler(ws))
}
