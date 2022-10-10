package dblayer

import (
	"backend/src/src/models"
	"github.com/jinzhu/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetPromos() (products models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(firstname, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error
}
