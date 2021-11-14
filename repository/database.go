package repository

import (
	"go-micro/tools"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	sqlDB, err := db.DB()

	if err != nil {
		sqlDB.SetMaxIdleConns(4)
		sqlDB.SetMaxOpenConns(20)
	} else {
		tools.GetLogger().Println("Cant set pool configuration")
	}

	db = database
}

func getDB() *gorm.DB {
	return db
}
