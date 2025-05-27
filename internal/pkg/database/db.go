package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)




func Init(dsn string) *gorm.DB {
	var err error

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	
	log.Println("Database connection established successfully")
	return db
}
