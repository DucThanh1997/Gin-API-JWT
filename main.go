package main

import (
	"Gin-API-JWT/Config"
	"Gin-API-JWT/Models"
	"Gin-API-JWT/Routers"
	"fmt"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", "root:thanh1997@tcp(127.0.0.1:3306)/hoc_sinh_go?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Student{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
