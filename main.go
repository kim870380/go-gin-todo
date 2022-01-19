package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo/config"
	"todo/models"
	"todo/routes"
)

var err error

func main() {
	dsn := config.DbURL(config.BuildDBConfig())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("status", err)
	}

	config.DB = db

	config.DB.AutoMigrate(&models.Todo{})

	//db.Create(&models.Todo{Title: "test", Description: "test"})

	r := routes.SetupRouter()

	r.Run()
}
