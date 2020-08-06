package main

import (
	config "Todo_demoapp/Todo_demoapp/Config"
	models "Todo_demoapp/Todo_demoapp/Models"
	routes "Todo_demoapp/Todo_demoapp/Routes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	// Creating a connection to the database
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	//setup routes ...
	r := routes.SetupRouter()

	// running
	r.Run()
}
