//package main
//
//import (
//	"gin-message-board/controllers"
//	"gin-message-board/tool"
//
//	"github.com/gin-gonic/gin"
//)
//
//func main() {
//	tool.SqlEngine()
//
//	engin := gin.Default()
//
//	registRouter(engin)
//
//	engin.Run()
//}
//
////调用controller
//func registRouter(engine *gin.Engine) {
//	new(controllers.HelloController).Router(engine)
//	new(controllers.UserController).Router(engine)
//	new(controllers.MessageController).Router(engine)
//}

package main

import (
	"gin-message-board/config"
	"gin-message-board/database"
	"gin-message-board/routers"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	config.Init()
	database.Init()
	var router = routers.InitializeRoutes()

	// 启动服务
	router.Run()

}
