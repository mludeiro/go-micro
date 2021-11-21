package database

import (
	"go-micro/entity"
	"go-micro/tools"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var config = &gorm.Config{
	Logger: logger.New(
		tools.GetLogger(), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		}),
}

func InitializePostgress() {

	dsn := "host=localhost user=postgres password=postgres dbname=go_micro port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), config)

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db = database
}

func InitializeSqlite() {

	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), config)

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db = database
}

func Migrate() {
	GetDB().AutoMigrate(&entity.ArticleType{}, &entity.Article{}, &entity.Client{}, &entity.Invoice{}, &entity.InvoiceLine{})
}

func CreateSampleData() {
	shoes := entity.ArticleType{Name: "Shoes"}
	pants := entity.ArticleType{Name: "Pants"}
	hats := entity.ArticleType{Name: "Hats"}

	GetDB().Create(&shoes).Create(&pants).Create(&hats)

	tennis := entity.Article{Name: "Tennis shoes", Price: 120, ArticleTypeID: shoes.ID}
	jeans := &entity.Article{Name: "Jeans", Price: 30, ArticleTypeID: pants.ID}

	GetDB().Create(&tennis).
		Create(&entity.Article{Name: "Running shoes", Price: 105, ArticleTypeID: shoes.ID}).
		Create(&entity.Article{Name: "Not to run shoes", Price: 88, ArticleTypeID: shoes.ID}).
		Create(&entity.Article{Name: "Jumping shoes", Price: 95, ArticleTypeID: shoes.ID}).
		Create(&entity.Article{Name: "Running but not Jumping shoes", Price: 67, ArticleTypeID: shoes.ID}).
		Create(&jeans).
		Create(&entity.Article{Name: "Oxford Jeans", Price: 60, ArticleTypeID: pants.ID}).
		Create(&entity.Article{Name: "Blue Jeans", Price: 40, ArticleTypeID: pants.ID}).
		Create(&entity.Article{Name: "Orange Jeans", Price: 40, ArticleTypeID: pants.ID}).
		Create(&entity.Article{Name: "All Colors Jeans", Price: 50, ArticleTypeID: pants.ID}).
		Create(&entity.Article{Name: "Cant belive its a jean", Price: 5, ArticleTypeID: pants.ID})

	carlos := entity.Client{Name: "Carlos", Address: "Siempreviva 123"}
	laura := entity.Client{Name: "Laura", Address: "Siempreviva 321"}
	pedro := entity.Client{Name: "Pedro", Address: "Calle Falsa 123"}

	GetDB().Create(&carlos).Create(&laura).Create(&pedro)

	invoice := entity.Invoice{ClientID: laura.ID, Amount: 0, Closed: false}

	GetDB().Create(&invoice)

	GetDB().Create(&entity.InvoiceLine{InvoiceID: invoice.ID, ArticleID: tennis.ID, Quantity: 1})
	GetDB().Create(&entity.InvoiceLine{InvoiceID: invoice.ID, ArticleID: jeans.ID, Quantity: 2})

	// We dont sell hats
}

func GetDB() *gorm.DB {
	return db
}
