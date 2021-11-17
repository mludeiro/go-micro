package repository

import (
	"go-micro/tools"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var dbLogger = logger.New(
	tools.GetLogger(), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		Colorful:                  true,        // Disable color
	},
)

func Initialize(onMemory bool) {
	dsn := "host=localhost user=postgres password=postgres dbname=go_micro port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	config := &gorm.Config{
		Logger: dbLogger,
	}

	var database *gorm.DB
	var err error

	if onMemory {
		database, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), config)
	} else {
		database, err = gorm.Open(postgres.Open(dsn), config)
	}

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db = database
}

func Migrate() {
	getDB().AutoMigrate(&ArticleType{}, &Article{})
}

func CreateSampleData() {
	shoes, _ := AddArticleType(&ArticleType{Name: "Shoes"})
	pants, _ := AddArticleType(&ArticleType{Name: "Pants"})
	AddArticleType(&ArticleType{Name: "Hats"})

	AddArticle(&Article{Name: "Tennis shoes", ArticleTypeID: shoes.ID})
	AddArticle(&Article{Name: "Running shoes", ArticleTypeID: shoes.ID})
	AddArticle(&Article{Name: "Not to run shoes", ArticleTypeID: shoes.ID})
	AddArticle(&Article{Name: "Jumping shoes", ArticleTypeID: shoes.ID})
	AddArticle(&Article{Name: "Running but not Jumping shoes", ArticleTypeID: shoes.ID})
	AddArticle(&Article{Name: "Jeans", ArticleTypeID: pants.ID})
	AddArticle(&Article{Name: "Oxford Jeans", ArticleTypeID: pants.ID})
	AddArticle(&Article{Name: "Blue Jeans", ArticleTypeID: pants.ID})
	AddArticle(&Article{Name: "Orange Jeans", ArticleTypeID: pants.ID})
	AddArticle(&Article{Name: "All Colors Jeans", ArticleTypeID: pants.ID})
	AddArticle(&Article{Name: "Cant belive its a jean", ArticleTypeID: pants.ID})

	// We dont sell hats
}

func getDB() *gorm.DB {
	return db
}
