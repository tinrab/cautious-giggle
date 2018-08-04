package domain

type Order struct {
  Entity
  Owner    Account
  Products []OrderedProduct
}

type OrderedProduct struct {
  Product
  Quantity uint32
}

func NewOrder(owner Account) Order {
  return Order{
    Entity:   newEntity(),
    Owner:    owner,
    Products: []OrderedProduct{},
  }
}

func (o *Order) AddProduct(p Product, quantity uint32) *Order {
  o.Products = append(o.Products, OrderedProduct{
    Product:  p,
    Quantity: quantity,
  })
  return o
}

func (o Order) TotalPrice() float64 {
  total := 0.0
  for _, p := range o.Products {
    total += p.Price
  }
  return total
}
