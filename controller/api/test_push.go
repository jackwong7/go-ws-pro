package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ws-pro/response_json"
	"ws-pro/ws"
)

func Push(c *gin.Context) {
	msg := ws.BroadCastMessageData{
		[]byte("hello every one"),
	}
	ws.WebsocketManager.BroadCastMessage <- &msg
	c.JSON(http.StatusOK, response_json.OK)
}
