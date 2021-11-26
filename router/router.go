package router

import (
	"github.com/gin-gonic/gin"
	"github.com/xhyonline/christmas_play_room/controllers"
)

// InitRouter 初始化路由
func InitRouter(engine *gin.Engine) {
	engine.GET("/index", controllers.StaticIndex)
}
