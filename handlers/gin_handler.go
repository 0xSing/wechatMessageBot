package handlers

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"runtime"
	"strategyBot/pkg"
)

type (
	SendWechatTextResp struct {
		Text string `json:"text"`
	}
)

func SendWechatText(c *gin.Context) {
	var ar SendWechatTextResp
	if err := c.ShouldBindJSON(&ar); err != nil {
		resp := pkg.MakeResp(pkg.ParamsError, nil)
		c.JSON(resp.HttpCode, resp)
		return
	}

	if err := SendTexts(ar.Text); err != nil {
		resp := pkg.MakeResp(pkg.WechatSendTextErr, nil)
		c.JSON(resp.HttpCode, resp)
		return
	}
	resp := pkg.MakeResp(pkg.Success, nil)
	c.JSON(resp.HttpCode, resp)
	return
}

func SendWechatImages(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		resp := pkg.MakeResp(pkg.ParamsError, nil)
		c.JSON(resp.HttpCode, resp)
		return
	}
	file := bytes.NewReader(data)
	if err := SendImages(file); err != nil {
		resp := pkg.MakeResp(pkg.WechatSendImagesErr, nil)
		c.JSON(resp.HttpCode, resp)
		return
	}
	resp := pkg.MakeResp(pkg.Success, nil)
	c.JSON(resp.HttpCode, resp)
	return
}

func CheckMemory(c *gin.Context) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("程序占用内存：%v bytes\n", m.Alloc)
	fmt.Printf("程序分配的总内存：%v bytes\n", m.TotalAlloc)
	fmt.Printf("程序已释放的内存：%v bytes\n", m.HeapReleased)
	resp := pkg.MakeResp(pkg.Success, nil)
	c.JSON(resp.HttpCode, resp)
	return
}
