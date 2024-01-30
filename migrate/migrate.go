package main

import (
	"github.com/kylerequez/go-gin-api/initializers"
	"github.com/kylerequez/go-gin-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DropUsersTable()
	initializers.DB.AutoMigrate(&models.User{})
}
