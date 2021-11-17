package repository

import (
	"go-micro/tools"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func Initialize() {
	newLogger := logger.New(
		tools.GetLogger(), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	dsn := "host=localhost user=postgres password=postgres dbname=go_micro port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	// sqlDB, err := db.DB()

	// if err != nil {
	// 	sqlDB.SetMaxIdleConns(4)
	// 	sqlDB.SetMaxOpenConns(20)
	// } else {
	// 	tools.GetLogger().Println("Cant set pool configuration")
	// }

	db = database
}

func getDB() *gorm.DB {
	return db
}
