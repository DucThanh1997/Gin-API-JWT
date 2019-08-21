package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
  )

type Student struct {
	Id int `gorm:"type:int;auto_increment;primary_key"`
	Name string `gorm:"type:Varchar(70);not null"`
	Age int `gorm:"type:int;not null"`
}

func Database() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:thanh1997@tcp(127.0.0.1:3306)/hoc_sinh_go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	
	db.AutoMigrate(&Student{})

	return db, nil
}

func Create(name1 string, age1 int) string {
	db, _ := Database()
	err1 := db.Save(&Student{Id: 1, Name: name1, Age: age1})
	fmt.Println(name1, age1)
	fmt.Println(err1)
	db.Close()
	return "okke"
}

// lấy hết học sinh
func GetAllStudent(b *[]Student) (err error) {
	db, _ := Database()
	if err = db.Find(b).Error; err != nil {
		return err
	}
	return nil
}