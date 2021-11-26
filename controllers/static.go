package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xhyonline/christmas_play_room/services"
	"net/http"
)

// StaticIndex 静态首页
func StaticIndex(c *gin.Context) {
	c.String(http.StatusOK, services.GetStaticHTML("index.html"))
}
