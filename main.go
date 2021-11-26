package main

import (
	"github.com/xhyonline/christmas_play_room/services"
	"net/http"

	"github.com/xhyonline/christmas_play_room/configs"
	"github.com/xhyonline/christmas_play_room/internal"
	"github.com/xhyonline/christmas_play_room/middleware"
	"github.com/xhyonline/christmas_play_room/router"

	"github.com/xhyonline/xutil/sig"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	// 初始化配置
	configs.Init(configs.WithMySQL())
	// 初始化 mysql 、redis 等服务组件
	//component.Init(component.RegisterMySQL())
	// 中间件
	g.Use(middleware.Cors(), middleware.Static())
	// 静态文件全部打包
	g.StaticFS("/static", http.FS(services.StaticFS))
	// 初始化路由
	router.InitRouter(g)
	// 启动 HTTP 服务
	httpServer := &internal.HTTPServer{Server: &http.Server{Addr: ":8080", Handler: g}}
	go httpServer.Run()
	// 注册优雅退出
	ctx := sig.Get().RegisterClose(httpServer)
	<-ctx.Done()
}
