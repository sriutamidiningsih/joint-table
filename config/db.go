package config

import (
	"dummy/utility"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Database() *gorm.DB {
	config, _ := utility.LoadConfig(".")
	db, err := gorm.Open(postgres.Open(config.DBSOURCE), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}
	fmt.Println("database connected")
	return db
}
