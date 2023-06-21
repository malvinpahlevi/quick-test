package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Product struct {
	ID          uint64 `gorm:"type:int;auto_increment;primary_key"`
	Price       uint64 `gorm:"type:int"`
	Description string `gorm:"type:varchar(40)"`
	Image       string
	Category    Category `gorm:"foreignkey:CategoryID;references:ID"`
	CategoryID  uint64
}

type Category struct {
	ID   uint64 `gorm:"type:int;auto_increment;primary_key"`
	Name string `gorm:"type:varchar(40)"`
}

type Customer struct {
	ID       uint64 `gorm:"type:int;auto_increment;primary_key"`
	Name     string `gorm:"type:varchar(40)"`
	Email    string `gorm:"type:varchar(40)"`
	Password string `gorm:"type:varchar(40)"`
}

type Order struct {
	ID      uint64 `gorm:"type:int;auto_increment;primary_key"`
	Product uint64
}

type OrderDetails struct {
	Quantity  uint64  `gorm:"type:int"`
	Order     Order   `gorm:"foreignkey:OrderID;references:ID"`
	OrderID   uint64  `gorm:"type:int;auto_increment;primary_key"`
	Product   Product `gorm:"foreignkey:ProductID;references:ID"`
	ProductID uint64  `gorm:"type:int"`
}

type Discount struct {
	ID    uint64 `gorm:"type:int;auto_increment;primary_key"`
	Name  string `gorm:"type:varchar(40)"`
	Rate  uint64 `gorm:"type:int"`
	Terms string `gorm:"type:varchar(40)"`
}

func main() {
	dsn := "root:admin@tcp(localhost:3306)/indochat_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connect db : %v", err)
		panic(err)
	}

	log.Printf("Success connect to database!")

	db.AutoMigrate(
		&Category{},
		&Product{},
		&Customer{},
		&Order{},
		&OrderDetails{},
		&Discount{})
	log.Printf("Success Automigrate!")
}
