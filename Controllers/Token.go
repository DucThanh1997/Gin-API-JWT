package Controllers

import (
	"Gin-API-JWT/ApiHelpers"
	"Gin-API-JWT/Models"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
	"strings"
	// "go-jwt/constants"
)
var jwtKey = []byte("đố em biết anh đang nghĩ gì")

func Signin(c *gin.Context) {
	var student Models.Student
	var student2 Models.Student
	
	// id := 1
	c.BindJSON(&student)
	fmt.Println(student.Username)
	err := Models.GetOneStudentByName(&student2, student.Username)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, student)
	}

	if student2.Password != student.Password {
		c.JSON(401, gin.H{
			"error": "sai mật khẩu",
		})
	} 
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := Models.Claims{
		Username: student.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	fmt.Println(tokenString)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		c.JSON(401, gin.H{
			"messages": "error",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":   "Signed in",
		"token": tokenString,
	})
}


func Refresh_token(c *gin.Context) {
	token := c.MustGet("token").(string)
	fmt.Println("token: ", token)
	claims := &Models.Claims{}
	parseToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	fmt.Println("jwtKey: ", jwtKey)
	fmt.Println("parseToken: ", parseToken)
	if !parseToken.Valid {
		c.JSON(401, gin.H{
			"msg": "Token is valid",
		})
		return
	}

	if err != nil {
		c.JSON(500, gin.H{
			"msg": "lỗi",
		})
		return
	}

	fmt.Println("claims.ExpiresAt: ", claims.ExpiresAt)
	
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < 10*time.Second {
		c.JSON(400, gin.H{
			"msg": "Token expired", 
		})
		return
	}
	expirationTime := time.Now().Add(24*time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	newToken :=jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := newToken.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{
			"msg": "error unknown",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":   "Refresh token",
		"token": tokenString,
	})
	
}


// jwt required
func JwtRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		token := ""
		arr := strings.Split(header, " ")
		if len(arr) > 1 {
			token = arr[1]
		} else {
			token = header 
		}
		if len(token) < 1 {
			c.JSON(401, gin.H{
				"error": "Unauthorized",
			})
		}
		claims := &Models.Claims{}
		parseToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if !parseToken.Valid || err != nil {
			c.JSON(500, gin.H{
				"message": "invalid token",
			})
			c.Abort()
			return 
		}
		
		var user Models.Student
		err = Models.GetOneStudentByName(&user, claims.Username)
		if err != nil {
			c.JSON(401, gin.H{
				"msg": "User not exist",
			})
			c.Abort()
			return
		}
		c.Set("token", token)
		c.Set("claims", claims)
		c.Set("userID", claims.Id)
		c.Next()
	}
}
