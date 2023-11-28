package main

import (
	"github.com/eatmoreapple/openwechat"
	"log"
	"os"
	"strategyBot/config"
	"strategyBot/handlers"
	"strategyBot/router"
)

func main() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 创建热存储容器对象
	reloadStorage := openwechat.NewJsonFileHotReloadStorage("storage.json")
	// 执行热登录
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Printf("login error: %v \n", err)
			return
		}
	}

	config.LoadConfig()

	config.UserBot, err = bot.GetCurrentUser()
	config.Config.RobotId = config.UserBot.ID()
	config.Friends, err = config.UserBot.Friends()

	// 初始化gin router
	err = router.Init(config.Config.HttpPort, "debug")
	if err != nil {
		log.Printf("start server failed")
		os.Exit(1)
	}
}
