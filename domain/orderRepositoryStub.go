package domain

type OrderRepositoryStub struct {
	orders []Order
}

func (s OrderRepositoryStub) GetAll(tenantId int, userName string) ([]Order, error) {
	return s.orders, nil
}

func NewOrderRepositoryStub() OrderRepositoryStub {
	orders := []Order{
		{OrderId: 123, Products: []Product{
			{ProductId: 100, ProductName: "Shirt", Price: 20.5, Quantity: 2},
			{ProductId: 120, ProductName: "T-Shirt", Price: 25.0, Quantity: 1},
		}},
	}

	return OrderRepositoryStub{orders: orders}
}
