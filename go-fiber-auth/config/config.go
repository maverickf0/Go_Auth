package config

import (
	"go-fiber-auth/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ConnectDB() {
	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	dsn := "host=" + host + " port=" + port + " user=" + user + " dbname=" + database + " password=" + password + " sslmode=disable"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
	}

	log.Println("Database connected successfully!")
}
