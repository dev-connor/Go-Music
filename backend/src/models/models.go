package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImagAlt     string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"` // sql.NullFloat64
	ProductName string  `json:"productname" gorm:"column:productname"`
	Description string
}

func (Product) TableName() string {
	return "products"
}

type Customer struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	LoggedIn  bool   `json:"loggedin"`
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `gorm:"column:customer_id"`
	ProductID    int       `gorm:"column:product_id"`
	Price        float64   `json:"sell_price" gorm:"column:price"`
	PurchaseDate time.Time `json:"purchase_date" gorm:"column:purchase_date"`
}

func (Order) TableName() string {
	return "orders"
}
