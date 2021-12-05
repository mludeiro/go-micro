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

type Database struct {
	gormDB *gorm.DB
}

func (db *Database) Migrate() *Database {
	db.GetDB().AutoMigrate(&entity.ArticleType{}, &entity.Article{}, &entity.Client{}, &entity.Invoice{}, &entity.InvoiceLine{})
	return db
}

func (db *Database) CreateSampleData() *Database {
	shoes := entity.ArticleType{Name: "Shoes"}
	pants := entity.ArticleType{Name: "Pants"}
	hats := entity.ArticleType{Name: "Hats"}

	db.GetDB().Create(&shoes).Create(&pants).Create(&hats)

	tennis := entity.Article{Name: "Tennis", Price: 120, ArticleTypeID: shoes.ID}
	jeans := &entity.Article{Name: "Jeans", Price: 30, ArticleTypeID: pants.ID}

	db.GetDB().Create(&tennis).
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

	db.GetDB().Create(&carlos).Create(&laura).Create(&pedro)

	invoice := entity.Invoice{ClientID: laura.ID, Amount: 0, Closed: false}

	db.GetDB().Create(&invoice)

	db.GetDB().Create(&entity.InvoiceLine{InvoiceID: invoice.ID, ArticleID: tennis.ID, Quantity: 1})
	db.GetDB().Create(&entity.InvoiceLine{InvoiceID: invoice.ID, ArticleID: jeans.ID, Quantity: 2})

	// We dont sell hats

	return db
}

func (db *Database) InitializePostgress() *Database {
	dsn := "host=localhost user=postgres password=postgres dbname=go_micro port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err := gorm.Open(postgres.Open(dsn), createGormConfig())

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db.gormDB = DB
	return db
}

func (db *Database) InitializeSqlite() *Database {

	DB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), createGormConfig())

	if err != nil {
		log.Fatal("Cannot initialize database")
	}

	db.gormDB = DB
	return db
}

func createGormConfig() *gorm.Config {
	return &gorm.Config{
		Logger: logger.New(
			tools.GetLogger(), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			}),
	}
}

func (db *Database) GetDB() *gorm.DB {
	return db.gormDB
}

func (db *Database) GetQueryDB(query Query) *gorm.DB {
	tx := db.GetDB()

	for _, fetch := range query.Fetchs {
		tx = tx.Preload(fetch)
	}

	for _, cond := range query.Conditions {
		switch cond.Comparator {
		case "eq":
			tx = tx.Where(cond.Field, cond.Value)
		}
	}

	for _, order := range query.OrderBy {
		tx = tx.Order(order)
	}
	return tx
}
