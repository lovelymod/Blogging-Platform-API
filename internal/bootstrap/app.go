package bootstrap

import (
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	DB *gorm.DB
}

func App() Application {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &Application{}
	app.DB = SetupDatabase() // เรียกฟังก์ชันจาก database.go
	// app.Env = NewEnv()
	return *app
}
