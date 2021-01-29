package controllers

import (
	"gin-message-board/database"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 生成随机的16字符字符串作为会话标记(生产环境不应使用这种方式)
func GenerateSessionToken() string {

	return strconv.FormatInt(rand.Int63(), 16)
}

// 展示注册页面
func ShowRegistrationPage(c *gin.Context) {
	render(c, gin.H{"title": "注册"}, "register.html")
}

// 注册
func Register(c *gin.Context) {
	// 获取"username"和"password"
	username := c.PostForm("username")
	password := c.PostForm("password")

	if err := database.RegisterNewUser(username, password); err == nil {
		// 如果创建了用户，需要在cookie中设置token，然后登录
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{"title": "成功注册，登录成功"}, "login-successful.html")

	} else {
		// 如果用户名或密码不合法展示错误再登录界面
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "注册失败",
			"ErrorMessage": err.Error()})

	}
}

// 展示登录界面
func ShowLoginPage(c *gin.Context) {
	render(c, gin.H{
		"title": "登录",
	}, "login.html")
}

// 用户登录
func PerformLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if valid, err := database.IsUserValid(username, password); err == nil && valid {
		token := GenerateSessionToken()
		c.SetCookie("token", token, 3600, "", "", false, true)

		render(c, gin.H{
			"title": "成功登录"}, "login-successful.html")

	} else {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "登录失败",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

// 用户登出
func Logout(c *gin.Context) {
	// 设置Cookie token令牌
	c.SetCookie("token", "", -1, "", "", false, true)
	// 重定向到首页
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
