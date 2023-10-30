package router

import (
	"fmt"
	"log"
	"strategyBot/middleware"

	"github.com/gin-gonic/gin"
)

func Init(port int, mode string) error {
	log.Println("starting server...")
	switch mode {
	case gin.DebugMode, gin.ReleaseMode, gin.TestMode:
	default:
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	router := gin.Default()
	router.Use(middleware.Cors()) // 跨域处理

	api := router.Group("/api")
	{
		initMsgRouter(api)
	}

	//address, err := middleware.GetIPAddress()
	//if err != nil {
	//	log.Println("Get ip Address fail: " + err.Error())
	//}

	log.Println("server started success...")
	return router.Run(fmt.Sprintf(":%d", port))
}
