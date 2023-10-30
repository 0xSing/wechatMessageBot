package handlers

import (
	"github.com/eatmoreapple/openwechat"
	"io"
	"log"
	"strategyBot/config"
	"strings"
)

var _ MessageHandlerInterface = (*UserMessageHandler)(nil)

// UserMessageHandler 私聊消息处理
type UserMessageHandler struct {
}

func (g *UserMessageHandler) handle(msg *openwechat.Message) error {
	sender, _ := msg.Sender()
	if !strings.EqualFold(sender.ID(), config.Config.RobotId) || strings.EqualFold(msg.Content, "stop syn check") {
		return g.ReplyText(msg)
	}
	return nil
}

func (g *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	// 接收私聊消息
	sender, err := msg.Sender()
	log.Printf("Received User %v Text Msg : %v", sender.NickName, msg.Content)
	context := "嘘，我在看盘！" + openwechat.Emoji.Shhh
	msg.ReplyText(context)
	log.Printf("Send User %v Text Msg : %v", sender.NickName, context)
	return err
}

// NewUserMessageHandler 创建私聊处理器
func NewUserMessageHandler() MessageHandlerInterface {
	return &UserMessageHandler{}
}

func SendTexts(context string) error {
	friends := config.Friends.SearchByID(config.Config.UserId)
	if len(friends) == 1 {
		log.Printf("Send User %v Text Msg : %v", friends[0].NickName, context)
		_, err := config.UserBot.SendTextToFriend(friends[0], context)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
	} else {
		log.Printf("Send User err: don't have friend with ID: ", config.Config.UserId)
	}
	return nil
}

func SendImages(file io.Reader) error {
	friends := config.Friends.SearchByID(config.Config.UserId)
	if len(friends) == 1 {
		log.Printf("Send User %v Text Msg : %v", friends[0].NickName, "images")
		_, err := config.UserBot.SendImageToFriend(friends[0], file)
		if err != nil {
			log.Fatal(err.Error())
			return err
		}
	} else {
		log.Printf("Send User err: don't have friend with ID: ", config.Config.UserId)
	}
	go stopSynCheck()
	return nil
}

func stopSynCheck() {
	//time.Sleep(time.Second * 6)
	helper := config.UserBot.FileHelper()
	log.Printf("Send User %v Text Msg : %v", helper.NickName, "stop syn check")
	_, err := config.UserBot.SendTextToFriend(helper, "stop syn check")
	if err != nil {
		log.Fatal(err.Error())
	}

}
