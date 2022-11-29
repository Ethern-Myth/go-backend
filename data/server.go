package data

import (
	"backend_sqlite/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open("sqlite3","user.db")

	if err != nil{
		panic("Failed to connect to database")
	}
	database.AutoMigrate(&models.User{})
	DB = database
}