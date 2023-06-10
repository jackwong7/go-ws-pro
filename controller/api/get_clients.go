package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ws-pro/response_json"
	"ws-pro/ws"
)

func GetClients(c *gin.Context) {
	c.JSON(http.StatusOK, response_json.OK.WithData(gin.H{"online": ws.WebsocketManager.GetClientCount()}))
}
