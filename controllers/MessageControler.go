package controllers

import (
	"fmt"

	"gin-message-board/service"
	"gin-message-board/tool"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
}

func (mc *MessageController) Router(engine *gin.Engine) {

	engine.POST("/messages", mc.SendMsg)
	engine.POST("/anonymous_messages", mc.anonymousMsg)
	engine.DELETE("/messages/:id", AdminMiddleWare(), mc.deleteMsg)
	engine.GET("/messages", mc.listMsgs)
	engine.GET("/messages/:id", mc.getOneMsg)
	engine.POST("/messages/:id/comment", mc.SendComment)
	engine.GET("/messages/:id/comment_list", mc.listComment)
}

//匿名留言
func (mc *MessageController) anonymousMsg(ctx *gin.Context) {
	username := tool.CheckLogin(ctx)
	if username == "" {
		tool.PrintInfo(ctx, "先登录再进行操作 ")
		return
	}

	message := ctx.PostForm("message")
	ms := service.MessageService{}
	err := ms.SendMsg(message, "anonymous")
	if err != nil {
		fmt.Println(err)
		return
	}

	tool.PrintInfo(ctx, "发表留言成功 ")
}

//单独列出一个留言的信息
func (mc *MessageController) getOneMsg(ctx *gin.Context) {
	id := ctx.Param("id")
	ms := service.MessageService{}
	err := ms.GetOneMsg(ctx, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//列出一个留言，及其下属的所有评论及回复
func (mc *MessageController) listComment(ctx *gin.Context) {
	//从path中获取pid
	pid := ctx.Param("id")
	//先输出root的信息
	ms := service.MessageService{}
	err := ms.GetOneMsg(ctx, pid)
	if err != nil {
		fmt.Println(err)
		return
	}

	Info, err := ms.GetInfo(pid)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Info) //test

	err = ms.TaoWa(ctx, Info)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//对一个留言或者评论，发送一条评论
func (mc *MessageController) SendComment(ctx *gin.Context) {
	//必须先登录才能发送
	username := tool.CheckLogin(ctx)
	if username == "" {
		tool.PrintInfo(ctx, "先登录再进行操作")
		return
	}
	//获取相关信息
	pid := ctx.Param("id")
	message := ctx.PostForm("message")

	ms := service.MessageService{}
	err := ms.SendComment(pid, username, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	tool.PrintInfo(ctx, "发送评论成功！")
}

//列出所有留言(不列出评论)
func (mc *MessageController) listMsgs(ctx *gin.Context) {
	ms := service.MessageService{}
	err := ms.ListMsg(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//删除一条留言
func (mc *MessageController) deleteMsg(ctx *gin.Context) {
	id := ctx.Param("id")
	ms := service.MessageService{}
	err := ms.DeleteMsg(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	tool.PrintInfo(ctx, "删除成功！")
}

//新建一条留言
func (mc *MessageController) SendMsg(ctx *gin.Context) {
	username := tool.CheckLogin(ctx)
	if username == "" {
		tool.PrintInfo(ctx, "先登录在进行操作 ")
		return
	}

	message := ctx.PostForm("message")
	ms := service.MessageService{}
	err := ms.SendMsg(message, username)
	if err != nil {
		fmt.Println(err)
		return
	}

	tool.PrintInfo(ctx, "发表留言成功 ")
}

//管理员权限中间件
func AdminMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("isLogin")
		if err == nil {
			username := cookie.Value
			if username == "wmf" {
				ctx.Next()
				return
			}
		}
		tool.PrintInfo(ctx, "你不是管理员！")
		ctx.Abort()
		return
	}
}
