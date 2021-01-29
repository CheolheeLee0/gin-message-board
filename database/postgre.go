package database

import (
	"errors"
	"fmt"
	"gin-message-board/models"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		viper.GetString("database.server"),
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.dbname"),
		viper.GetInt("database.ports"),
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("Fatal database error: \n", err))
	}

	db.AutoMigrate(&models.Message{}, &models.User{})
}

func GetDB() *gorm.DB {
	return db
}

func GetAllMessages() ([]models.Message, error) {
	var messages []models.Message
	result := db.Find(&messages)
	return messages, result.Error
}

func GetMessageByID(id int) (models.Message, error) {
	var message models.Message
	result := db.Where("ID = ?", id).First(&message)
	return message, result.Error
}

func CreateNewMessage(title, content string) (models.Message, error) {
	message := models.Message{Title: title, Content: content}
	result := db.Create(&message)
	return message, result.Error
}

func RegisterNewUser(username, password string) error {
	if strings.TrimSpace(password) == "" {
		return errors.New("密码不能为空")
	} else if user_avial, err := IsUsernameAvailable(username); err == nil && !user_avial {
		return errors.New("用户名不可用")
	}
	user := models.User{Username: username, Password: password}
	result := db.Create(&user)
	return result.Error
}

func IsUsernameAvailable(username string) (bool, error) {
	var user []models.User
	result := db.Where("Username = ?", username).Find(&user)
	return result.RowsAffected == 0, result.Error
}

func IsUserValid(username, password string) (bool, error) {
	var user []models.User
	result := db.Where("Username = ? AND Password = ?", username, password).Find(&user)
	return result.RowsAffected == 1, result.Error
}
