package http_server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/panjf2000/ants/v2"
	"log"
	"net/http"
	"time"
	"ws-pro/middleware"
	"ws-pro/route"
	"ws-pro/ws"
)

var (
	srv    *http.Server
	wsExit context.CancelFunc
)

func WebsocketServerRun() {
	var ctx context.Context
	ctx, wsExit = context.WithCancel(context.Background())
	go ws.WebsocketManager.Start(ctx)
	for i := 0; i < 2; i++ {
		go ws.WebsocketManager.SendService(ctx)
		go ws.WebsocketManager.SendGroupService(ctx)
		go ws.WebsocketManager.SendAllService(ctx)
	}
}

func HttpServerRun() {
	ws.GetInstance().InitPool(3, ants.WithOptions(ants.Options{MaxBlockingTasks: 2, Nonblocking: false}))

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())

	// register route
	apiRouter := route.NewApiRouter()
	apiRouter.InitRouter(r)

	srv = &http.Server{
		Addr:         ":8899",
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("server listen err:%s", err)
	}
}
func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	log.Println("正在停止api服务")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	ws.GetInstance().Close()
	log.Println("api服务已停止")
}
func WebsocketServerStop() {
	log.Println("正在停止ws服务")
	wsExit()
	ws.WebsocketManager.Wg.Wait()
	log.Println("ws服务已停止")
}
