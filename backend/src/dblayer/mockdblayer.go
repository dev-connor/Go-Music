package dblayer

import "backend/src/src/models"

type MockDBLayer struct {
	err       error
	products  []models.Product
	customers []models.Customer
	orders    []models.Order
}
