package database

import (
	"main/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)
var DB *gorm.DB

func Connect(url_db string) {

	connection, errr := gorm.Open(mysql.Open(url_db), &gorm.Config{})

	if errr != nil {
		panic("could not connect to db")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
