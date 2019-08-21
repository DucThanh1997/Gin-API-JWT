package Controllers

import (
	"../ApiHelpers"
	"../Models"
	"github.com/gin-gonic/gin"
)

func ListStudent(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllStudent(&student)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, student)
	} else {
		ApiHelpers.RespondJSON(c, 200, student)
	}
}

func AddNewStudent(c *gin.Context) {
	var student Models.Student{name: c.PostForm("name"), age: c.PostForm("age")}
	c.BindJSON(&student)
	err := Models.AddNewStudent(&student)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, student)
	} else {
		ApiHelpers.RespondJSON(c, 200, student)
	}
}

func GetOneStudent(c *gin.Context) {
	id := c.Params.ByName("id")
	var student Models.Student
	err := Models.GetOneStudent(&student, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, student)
	} else {
		ApiHelpers.RespondJSON(c, 200, student)
	}
}

func PutOneStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.GetOneStudent(&student, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, student)
	}
	c.BindJSON(&student)
	err = Models.PutOneStudent(&student, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, student)
	} else {
		ApiHelpers.RespondJSON(c, 200, student)
	}
}

func DeleteStudent(c *gin.Context) {
	var student Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&student, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, student)
	} else {
		ApiHelpers.RespondJSON(c, 200, student)
	}
}
