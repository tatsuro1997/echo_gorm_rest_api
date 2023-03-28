package model

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	godotenv.Load(".env")

	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}

	DB.AutoMigrate(&User{}, &Post{})
}
