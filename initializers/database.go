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

	var host string = os.Getenv("DB_HOST")
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var name string = os.Getenv("DB_NAME")
	var port string = os.Getenv("DB_PORT")
	var mode string = os.Getenv("DB_SSLMODE")
	var timezone string = os.Getenv("DB_TIMEZONE")

	var dsnTemplate = os.Getenv("DB_URL")

	dsn := fmt.Sprintf(dsnTemplate, host, user, password, name, port, mode, timezone)
	fmt.Println("The connection string is:::-:::", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
		return
	}

	fmt.Println(":::-:::\tCONNECTION SUCCESSFUL\t:::-:::")
}

func DropUsersTable() {
	DB.Migrator().DropTable(&models.User{})
}
