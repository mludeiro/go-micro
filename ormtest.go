package main

// import (
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type Product struct {
// 	gorm.Model
// 	Code      string
// 	Price     uint
// 	CompanyID int
// 	company   *Company
// }

// type Company struct {
// 	gorm.Model
// 	ID       int
// 	Name     string
// 	products []*Product
// }

// type User struct {
// 	gorm.Model
// 	Username string
// 	Orders   []Order
// }

// type Order struct {
// 	gorm.Model
// 	UserID uint
// 	Price  float64
// 	User   User
// }

// func main() {
// 	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if db == nil {
// 		log.Fatal("no database")
// 	}

// 	err = db.AutoMigrate(&User{}, &Order{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//	var users = []User{}
// 	var orders = []Order{}

// 	// Preload Orders when find users
// 	//	db.Debug().Find(&users)

// 	//	db.Debug().Preload("Users").Preload("Orders").Find(&users, "users.id IN ?", []int{1, 2, 3, 4, 5})
// 	// SELECT * FROM users;
// 	// SELECT * FROM orders WHERE user_id IN (1,2,3,4);

// 	db.Debug().Preload("User").Find(&orders)
// 	// SELECT * FROM users;
// 	// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
// 	// SELECT * FROM profiles WHERE user_id IN (1,2,3,4); // has one
// 	// SELECT * FROM roles WHERE id IN (4,5,6); // belongs to

// 	log.Println(orders[0].User)
// 	log.Println(orders)

// 	//	db.Debug().Create(&User{Username: "carlos", Orders: []Order{{Price: 12}}})

// 	/*
// 		var product = Product{Code: "EEE"}
// 		var company Company
// 		db.Debug().Preload("company").Find(&product)

// 		// find product with integer primary key

// 		log.Println(product.Code)
// 		log.Println(company.Name)

// 		company = Company{Name: "AWS2"}
// 			db.Create(&company)

// 			db.Create(&Product{Code: "EEE", Price: 50, company: company})
// 	*/
// }
