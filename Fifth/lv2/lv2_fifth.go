package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var account map[string]string

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
		c.JSON(401, gin.H{"error": "你还没有登录"})
		//若验证不通过，不再调用后续函数
		c.Abort()
		return
	}
}
func main() {
	account = make(map[string]string, 100)
	r := gin.Default()
	//加载html文件
	r.LoadHTMLGlob("tamplates/*")
	//注册页面
	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", "")
	})
	r.POST("/registerOn", func(c *gin.Context) {
		name := c.PostForm("username")
		password := c.PostForm("password")
		//储存账号密码
		account[name] = password
		//注册完自动进入登录页
		c.Redirect(http.StatusMovedPermanently, "/login")
	})
	//登录页面
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", "")
	})
	//校验
	r.POST("/loginOn", func(c *gin.Context) {
		name := c.PostForm("username")
		password := c.PostForm("password")
		if account[name] == password {
			//设置cookie:可保持1h登录状态
			c.SetCookie("status", "on", 3600, "/", "localhost", false, true)
			c.Redirect(http.StatusMovedPermanently, "/home")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/wrong")
		}
	})
	//主页
	r.GET("/home", LgcMiddleWare(), func(c *gin.Context) {
		c.String(200, "Welcome to home!")
	})
	//不会弹窗只能用跳页面代替了
	r.GET("/wrong", func(c *gin.Context) {
		c.String(http.StatusOK, "用户名或密码错误")
	})
	r.Run(":1010")
}
