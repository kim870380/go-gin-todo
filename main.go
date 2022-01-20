package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"todo/config"
	"todo/models"
	"todo/routes"
	"xorm.io/xorm"
)

var err error

func main() {
	dsn := config.DbURL(config.BuildDBConfig())
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	config.DB, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("status", err)
	}

	/* grom
	config.DB.AutoMigrate(&models.Todo{})
	*/
	if err = config.DB.Sync2(new(models.Todo)); err != nil {
		fmt.Println("status", err)
	}

	r := routes.SetupRouter()

	r.Run()
}
