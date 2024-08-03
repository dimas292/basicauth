package database

import (
	"basic-auth/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	dsn := "root:@tcp(127.0.0.1:3306)/basic_auth?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("hello")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})

}
