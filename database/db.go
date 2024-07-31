package db

import (
	"github.com/piotroszko/backend-go/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func DbConnect() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DBConn = db
}

func DbAutoMigrate() {
	DBConn.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
}
