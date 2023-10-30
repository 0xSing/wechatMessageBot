package router

import (
	"github.com/gin-gonic/gin"
	"strategyBot/handlers"
	"strategyBot/middleware"
)

// To wechat server接口
func initMsgRouter(r *gin.RouterGroup) {
	msgApiRouter := r.Group("/msg", middleware.Auth())
	{
		msgApiRouter.POST("/sendText", handlers.SendWechatText)
		msgApiRouter.POST("/sendImages", handlers.SendWechatImages)
	}

	testApiRouter := r.Group("/check")
	{
		testApiRouter.GET("/memory", handlers.CheckMemory)
	}
}
