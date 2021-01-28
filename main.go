package main

import (
	"gin-message-board/controllers"
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
	new(controllers.HelloController).Router(engine)
	new(controllers.UserController).Router(engine)
	new(controllers.MessageController).Router(engine)
}
