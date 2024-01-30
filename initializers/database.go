package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/kylerequez/go-gin-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		return
	}

	fmt.Println("CONNECTION SUCCESSFUL:::-:::", &DB)
}

func DropUsersTable() {
	DB.Migrator().DropTable(&models.User{})
}
