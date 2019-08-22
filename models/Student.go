package Models

import (
	"Gin-API-JWT/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllStudent(b *[]Student) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewStudent(b *Student) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneStudent(b *Student, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneStudentByName(b *Student, name string) (err error,) {
	if err := Config.DB.Where("username = ?", name).First(b).Error; err != nil {
		
		return err
	}
	return nil
}

func PutOneStudent(b *Student, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteStudent(b *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
