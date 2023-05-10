package main

import (
	"github.com/eswulei/chatgpt-web/pkg/model"
)

func main() {
	DB := model.ConnectDB()
	DB.AutoMigrate(&model.User{}, &model.Message{})
}
