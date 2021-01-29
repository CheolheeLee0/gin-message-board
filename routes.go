package main

func initializeRoutes() {

	// 处理主页路由
	router.GET("/", showIndexPage)

	router.GET("/message/view/:message_id", getMessage)

}
