package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"strategyBot/pkg"
)

// 授权中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查客户端IP是否属于内网
		r := c.Request
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println("无法获取客户端IP地址", http.StatusInternalServerError)
			resp := pkg.MakeResp(pkg.NotFoundInternalIp, nil)
			c.JSON(resp.HttpCode, resp)
			c.Abort()
			return
		}

		// 判断IP是否属于内网
		isInternal := false
		parsedIP := net.ParseIP(ip)
		if parsedIP != nil {
			if parsedIP.IsLoopback() || parsedIP.IsPrivate() {
				isInternal = true
			}
		}

		// 如果不是内网IP，则返回错误
		if !isInternal {
			log.Println("禁止访问", http.StatusForbidden)
			resp := pkg.MakeResp(pkg.NotInternalIp, nil)
			c.JSON(resp.HttpCode, resp)
			c.Abort()
			return
		}

		// 处理请求
		c.Next()
	}
}
