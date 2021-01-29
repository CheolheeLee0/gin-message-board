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
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	// 在一开始就处理模板，这样就不必再从磁盘加载它们了。
	router.LoadHTMLGlob("templates/*")

	// 初始化路由
	initializeRoutes()

	// 启动服务
	router.Run()

}
