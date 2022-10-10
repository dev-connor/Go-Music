package dblayer

type DBLayer interface {
	GetAllProducts() ([]model.Product, error)
}
