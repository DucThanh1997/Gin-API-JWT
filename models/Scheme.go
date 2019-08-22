package Models

import (
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"
)
// var jwtKey = []byte("my_secret_key")

type Student struct {
	gorm.Model
	Username     string `json:"username"`
	Age   		 string `json:"age"`
	Password 	 string `json:"password"`

}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (b *Student) TableName() string {
	return "student"
}
