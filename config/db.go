package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	config := godotenv.Load(".env")
	dbms := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Jakarta"
	db, config := gorm.Open(postgres.Open(dbms), &gorm.Config{})
	if config != nil {
		log.Fatal("DB connection error")
	}
	fmt.Println("database connected")
	return db
}
