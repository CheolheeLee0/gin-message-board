package controllers

import (
	"fmt"

	"gin-message-board/service"
	"gin-message-board/tool"

	"github.com/gin-gonic/gin"

	"net/http"
)

type UserController struct {
}

// 用户路由
// TODO "Router" vs "UserRouter" ? 哪种可读性更强
func (uc *UserController) Router(engine *gin.Engine) {
	engine.POST("/register", uc.register)         //用户注册
	engine.POST("/login", uc.login)               //用户登录
	engine.POST("/change_password", uc.changePwd) //用户改密
	engine.GET("/logout", uc.logout)              //用户登出
}

//用户注册
func (uc *UserController) register(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	fmt.Println("RegisterUserInfo:", userName, pwd)

	us := new(service.UserService)

	flag := us.CheckUserAlive(userName)
	if flag == false {
		tool.PrintInfo(ctx, "该用户已经存在 ")
		return
	}

	ok := us.RegisteByPwd(userName, pwd)
	if ok == true {
		tool.PrintInfo(ctx, "注册成功 ")
		return
	}
	tool.PrintInfo(ctx, "注册失败 ")

}

//用户登录
func (uc *UserController) login(ctx *gin.Context) {
	value := tool.CheckLogin(ctx)
	if value != "" {
		tool.PrintInfo(ctx, "用户已登录 ")
		return
	}

	userName := ctx.PostForm("username")
	pwd := ctx.PostForm("password")
	fmt.Println("LoginUserInfo:", userName, pwd)

	us := new(service.UserService)

	flag := us.CheckUserAlive(userName)
	if flag == true {
		tool.PrintInfo(ctx, "该用户不存在 ")
		return
	}

	cookie := us.LoginByPwd(userName, pwd)
	if cookie == nil {
		tool.PrintInfo(ctx, "密码错误 ")
		return
	}

	http.SetCookie(ctx.Writer, cookie)
	tool.PrintInfo(ctx, "登录成功 ")
}

//退出登录
func (uc *UserController) logout(ctx *gin.Context) {
	value := tool.CheckLogin(ctx)
	if value == "" {
		tool.PrintInfo(ctx, "未登录 ")
		return
	}

	cookie, err := ctx.Request.Cookie("isLogin")
	if err != nil {
		tool.PrintInfo(ctx, "获取cookie失败")
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(ctx.Writer, cookie)

	tool.PrintInfo(ctx, "退出登录成功")
}

//修改密码
func (uc *UserController) changePwd(ctx *gin.Context) {
	//验证登录状态，只有登录才能修改密码
	username := tool.CheckLogin(ctx)
	newPwd := ctx.PostForm("newPwd")
	if username == "" {
		tool.PrintInfo(ctx, "请先登录 ")
		return
	}
	//service
	us := service.UserService{}
	err := us.ChangePwd(username, newPwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(username, "newPwd: ", newPwd)
	tool.PrintInfo(ctx, "修改密码成功")
}
