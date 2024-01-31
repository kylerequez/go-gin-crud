package main

import (
	"github.com/kylerequez/go-gin-api/controllers"
	"github.com/kylerequez/go-gin-api/initializers"
	"github.com/kylerequez/go-gin-api/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.New()

	// Users Endpoint
	v1 := r.Group("/api/v1/users")
	{
		v1.POST("/register", controllers.RegistrationHandler)
		v1.POST("/login", controllers.LoginHandler)
		v1.GET("/", controllers.GetUsers)
		v1.GET("/:id", controllers.GetUserById)
		v1.Use(middlewares.UserAuthMiddleware())
		v1.POST("/", controllers.CreateUser)
		v1.PATCH("/:id", controllers.PatchUpdateUser)
		v1.PUT("/:id", controllers.PutUpdateUser)
		v1.DELETE("/:id", controllers.DeleteUserById)
	}

	// Ping Endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	r.Run()
}
