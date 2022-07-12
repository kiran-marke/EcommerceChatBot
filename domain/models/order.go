package domain

type Product struct {
	ProductId   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type Order struct {
	OrderId  int       `json:"order_id"`
	Products []Product `json:"products"`
}

type OrderRepository interface {
	GetAll(int, string) ([]Order, error)
}
