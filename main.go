package main

import (
	"gin-message-board/controller"
	"gin-message-board/tool"

	"github.com/gin-gonic/gin"
)

func main() {
	tool.SqlEngine()

	engin := gin.Default()

	registRouter(engin)

	engin.Run()
}

//调用controller
func registRouter(engine *gin.Engine) {
	new(controller.HelloController).Router(engine)
	new(controller.UserController).Router(engine)
	new(controller.MessageController).Router(engine)
}
