package handler

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error loading .env file")
	}

	//open a db connection
	dba := os.Getenv("DB")
	driver := os.Getenv("DB_DRIVER")

	var err error
	db, err = gorm.Open(driver, dba)
	if err != nil {
		panic("failed to connect database")
	}

}
