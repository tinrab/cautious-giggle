package domain

type Product struct {
	Entity
	Name  string
	Price float64
}

func NewProduct(name string, price float64) Product {
	return Product{
		Entity: newEntity(),
		Name:   name,
		Price:  price,
	}
}
