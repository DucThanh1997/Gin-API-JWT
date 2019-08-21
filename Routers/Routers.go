package Routers

import (
	"../Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("student", Controllers.ListStudent)
		v1.POST("student", Controllers.AddNewStudent)
		v1.GET("student/:id", Controllers.GetOneStudent)
		v1.PUT("student/:id", Controllers.PutOneStudent)
		v1.DELETE("student/:id", Controllers.DeleteStudent)
	}

	return r
}
