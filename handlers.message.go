package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	messages := getAllMessages()

	// 调用上下文的HTML方法来呈现模板
	//c.HTML(
	//	// 设置HTTP状态为200 (OK)
	//	http.StatusOK,
	//	// 使用index.html模板
	//	"index.html",
	//	// 传递页面使用的数据
	//	gin.H{
	//		"title":   "主页",
	//		"payload": messages,
	//	},
	//)
	render(c, gin.H{
		"title":   "主页",
		"payload": messages}, "index.html")

}

func getMessage(c *gin.Context) {
	// 检查留言是否合法
	if messageID, err := strconv.Atoi(c.Param("message_id")); err == nil {
		// 检查留言是否存在
		if message, err := getMessageByID(messageID); err == nil {
			c.HTML(
				http.StatusOK,
				// 使用"message.html"模板
				"message.html",
				// 传递页面使用的数据
				gin.H{
					"title":   message.Title,
					"payload": message,
				},
			)

		} else {
			// 如果留言没找到，返回NotFound错误
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// 如果在URL中指定了无效的留言ID，则终止并显示错误
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// 响应JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// 响应XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// 默认响应HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
