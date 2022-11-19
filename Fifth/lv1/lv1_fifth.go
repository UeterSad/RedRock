package main

import (
	"github.com/gin-gonic/gin"
)

//通过 /login 与 /home 测试

// 2.设置一个中间件

func LgcMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取客户端cookie并校验
		if cookie, err := c.Cookie("status"); err == nil {
			if cookie == "on" {
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(401, gin.H{"error": "err"})
		//若验证不通过，不再调用后续函数
		c.Abort()
		return
	}
}
func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		//1.设置cookie
		//maxAge int 持续时间（单位为秒）
		//path cookie所在目录
		//domain 域名
		//secure 能否智能通过https访问
		//httponly 是否允许别人通过js获取自己的cookie
		c.SetCookie("status", "on", 60, "/", "localhost", false, true)
		c.String(200, "Login success!")
	})
	//3.设置home页
	r.GET("/home", LgcMiddleWare(), func(c *gin.Context) {
		c.String(200, "Welcome to home!")
	})
	r.Run(":1010")
}
