package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kylerequez/go-gin-api/controllers"
	"github.com/kylerequez/go-gin-api/initializers"
	"github.com/kylerequez/go-gin-api/services"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func customMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(":::-:::\tAUTHORIZATION MIDDLEWARE\t:::-:::")

		cookie, err := c.Cookie("go-gin-crud-token")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "The auth cookie does not exists",
			})
			c.Abort()
			return
		}

		isValidated, err := services.ValidateToken(cookie)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "There was an error in validating the token",
			})
			c.Abort()
			return
		}

		claims, ok := isValidated.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "There was an error in parsing the token",
			})
			return
		}

		exp := claims["expiration"].(float64)
		if int64(exp) < time.Now().Local().Unix() {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "The token has expired. Please login again.",
			})
			c.Abort()
			return
		}

		isAuthorized := claims["authorized"].(bool)
		if !isAuthorized {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "The token has expired. Please login again.",
			})
			c.Abort()
			return
		}
		fmt.Println(":::-:::\tAUTHORIZATION SUCCESSFUL\t:::-:::")
		c.Next()
	}
}

func main() {
	r := gin.New()

	// Users Endpoint
	v1 := r.Group("/api/v1")
	{
		v1.POST("/users/login", controllers.LoginHandler)
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/users/:id", controllers.GetUserById)
		v1.Use(customMiddleware())
		v1.POST("/users", controllers.CreateUser)
		v1.PATCH("/users/:id", controllers.PatchUpdateUser)
		v1.PUT("/users/:id", controllers.PutUpdateUser)
		v1.DELETE("/users/:id", controllers.DeleteUserById)
	}

	// Ping Endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	r.Run()
}
