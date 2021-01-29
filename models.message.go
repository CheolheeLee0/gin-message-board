package main

import "errors"

type message struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var messageList = []message{
	message{ID: 1, Title: "留言标题1", Content: "留言内容1"},
	message{ID: 2, Title: "留言标题2", Content: "留言内容2"},
}

// 返回留言列表
func getAllMessages() []message {
	return messageList
}

// 根据提供的ID获取一个留言
func getMessageByID(id int) (*message, error) {
	for _, m := range messageList {
		if m.ID == id {
			return &m, nil
		}
	}
	return nil, errors.New("Message not found")
}
