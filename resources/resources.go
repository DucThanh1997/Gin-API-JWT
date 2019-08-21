package resources

import "fmt"
import "github.com/gin-gonic/gin"
import "Gin-API-JWT/models"
import "strconv"

func HelloAPI(c *gin.Context) {
	fmt.Println("lalala")

}

func Create(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	ag3, _ := strconv.Atoi(age)
	try1 := models.Create(name, ag3)
	fmt.Println(try1)
	c.String(200, "okke")
}

