package domain

import domain "github.com/kiran-marke/ecommercechatbot/domain/models"

type OrderRepositoryStub struct {
	orders []domain.Order
}

func (s OrderRepositoryStub) GetAll(tenantId int, userName string) ([]domain.Order, error) {
	return s.orders, nil
}

func NewOrderRepositoryStub() OrderRepositoryStub {
	orders := []domain.Order{
		{OrderId: 123, Products: []domain.Product{
			{ProductId: 100, ProductName: "Shirt", Price: 20.5, Quantity: 2},
			{ProductId: 120, ProductName: "T-Shirt", Price: 25.0, Quantity: 1},
		}},
	}

	return OrderRepositoryStub{orders: orders}
}
