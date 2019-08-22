package Routers

import (
	"Gin-API-JWT/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("signin", Controllers.Signin)
	v1 := r.Group("/v1")
	v1.Use(Controllers.JwtRequired())
	{
		v1.GET("student", Controllers.ListStudent)
		v1.POST("student", Controllers.AddNewStudent)
		v1.GET("student/:id", Controllers.GetOneStudent)
		v1.PUT("student/:id", Controllers.PutOneStudent)
		v1.DELETE("student/:id", Controllers.DeleteStudent)
		v1.POST("token", Controllers.Refresh_token)
	}

	return r
}
