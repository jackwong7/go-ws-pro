package route

import (
	"github.com/gin-gonic/gin"
	"ws-pro/controller/api"
	"ws-pro/ws"
)

type ApiRouter struct{}

func NewApiRouter() *ApiRouter {
	return &ApiRouter{}
}

func (a *ApiRouter) InitRouter(router *gin.Engine) {
	router.GET("/getClients", api.GetClients)
	router.GET("/push", api.Push)
	router.GET("/ws", ws.WebsocketManager.WsClient)
}
