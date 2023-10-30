package config

import (
	"encoding/json"
	"github.com/eatmoreapple/openwechat"
	"log"
	"os"
	"strconv"
	"sync"
)

// Configuration 项目配置
type Configuration struct {
	// Owner ID
	UserId string `json:"user_id"`
	// 自动通过好友
	AutoPass bool `json:"auto_pass"`
	// Server port
	HttpPort int `json:"http_port"`
	// Robot ID
	RobotId string
}

var (
	UserBot *openwechat.Self
	Friends openwechat.Friends
	Config  *Configuration
)

var once sync.Once

// LoadConfig 加载配置
func LoadConfig() {
	once.Do(func() {
		// 从文件中读取
		Config = &Configuration{}
		f, err := os.Open("config.json")
		if err != nil {
			log.Fatalf("open config err: %v", err)
			return
		}
		//defer f.Close()
		encoder := json.NewDecoder(f)
		err = encoder.Decode(Config)
		if err != nil {
			log.Fatalf("decode config err: %v", err)
			return
		}

		// 如果环境变量有配置，读取环境变量
		AutoPass := os.Getenv("AutoPass")
		UserId := os.Getenv("UserId")
		if UserId != "" {
			Config.UserId = UserId
		}
		if AutoPass == "true" {
			Config.AutoPass = true
		}

		// Server configuration
		HttpPort := os.Getenv("HttpPort")
		if HttpPort != "" {
			atoi, _ := strconv.Atoi(HttpPort)
			println(atoi)
			Config.HttpPort = atoi
		}
	})
}
