package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	imageAlt    string  `json:"imgalt" // gorm:"column:imgalt"`
	Price       float64 `json:"price""`
	Promotion   float64 `json:"promotion"`   //sql.NullFloat64
	ProductName string  `json:"productname"` // gorm:"column:productname"
	Description string
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	gorm.Model
	Name      string  `json:"name" sql:"-"`
	FirstName string  `gorm:"column:firstname" json:"firstname"`
	LastName  string  `gorm:"column:lastname" json:"lastname"`
	Email     string  `gorm:"column:email" json:"email"`
	Pass      string  `json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders    []Order `json:"orders"`
}

func (Customer) TableName() string {
	return "customers"
}

type Order struct {
	gorm.Model
	Product      `sql:"-"`
	Customer     `sql:"-"`
	CustomerID   int       `json:"customer_id" gorm:"column:customer_id"`
	ProductID    int       `json:"product_id" gorm:"column:product_id"`
	Price        float64   `gorm:"column:price" json:"sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}
