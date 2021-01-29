package main

import (
	"gin-message-board/controllers"
)

func initializeRoutes() {

	// 处理主页路由
	router.GET("/", controllers.ShowIndexPage)

	router.GET("/message/view/:message_id", controllers.GetMessage)

}
