package controllers

import (
	"gin-message-board/database"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	//messages := models.GetAllMessages()
	messages, err := database.GetAllMessages()
	if err == nil {
		render(c, gin.H{
			"title":   "主页",
			"payload": messages}, "index.html")
	}

}

func GetMessage(c *gin.Context) {
	// 判断是否登录
	loggedInInterface, _ := c.Get("is_logged_in")
	// 检查留言是否合法
	if messageID, err := strconv.Atoi(c.Param("message_id")); err == nil {
		// 检查留言是否存在
		if message, err := database.GetMessageByID(messageID); err == nil {
			c.HTML(
				http.StatusOK,
				// 使用"message.html"模板
				"message.html",
				// 传递页面使用的数据
				gin.H{
					"title":        message.Title,
					"payload":      message,
					"is_logged_in": loggedInInterface.(bool),
				},
			)

		} else {
			// 如果留言没找到，返回NotFound错误
			_ = c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// 如果在URL中指定了无效的留言ID，则终止并显示错误
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)
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

// 留言创建页面
func ShowMessageCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create New Message Title"}, "create-message.html")
}

// 留言提交成功页面
func CreateMessage(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := database.CreateNewMessage(title, content); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
