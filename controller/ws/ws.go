package ws

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/gorilla/websocket"
//	"log"
//	"net/http"
//	"time"
//	ws_pool2 "ws-pro/ws_pool"
//)
//
//var Connections = make(map[*websocket.Conn]bool, 100)
//
//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//	// 解决跨域问题
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}
//
//func WebSocket(c *gin.Context) {
//	log.Println("received echo request.")
//	//服务升级，对于来到的http连接进行服务升级，升级到ws
//	cn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	Connections[cn] = true
//	fmt.Printf("%v\n", Connections)
//
//	cn.SetCloseHandler(func(code int, text string) error {
//		delete(Connections, cn)
//		message := websocket.FormatCloseMessage(code, "bye bye")
//		cn.WriteControl(websocket.CloseMessage, message, time.Now().Add(websocket.CloseMessage))
//		return nil
//	})
//	defer cn.Close()
//	if err != nil {
//		log.Println(err.Error())
//		return
//	}
//	for {
//		//messageType int, p []byte, err error
//		mt, message, err := cn.ReadMessage()
//		if err != nil {
//			log.Println("read:", err)
//			break
//		}
//		msg := ws_pool2.Message{
//			Conn:        cn,
//			MessageType: mt,
//			Message:     message,
//		}
//		go ws_pool2.GetInstance().NewTask(msg, ws_pool2.GetInstance().Process(msg))
//		log.Printf("recv: %s", message)
//
//		ws_pool2.GetInstance().Mux.Lock()
//		err = cn.WriteMessage(mt, []byte(fmt.Sprintf("received your message: %s", message)))
//		if err != nil {
//			log.Println("write:", err)
//			ws_pool2.GetInstance().Mux.Unlock()
//			break
//		}
//		ws_pool2.GetInstance().Mux.Unlock()
//	}
//}
