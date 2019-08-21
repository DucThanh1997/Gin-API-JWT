package Models

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Age   string `json:"age"`
}

func (b *Student) TableName() string {
	return "student"
}
