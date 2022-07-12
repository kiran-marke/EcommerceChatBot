package domain

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	domain "github.com/kiran-marke/ecommercechatbot/domain/models"
)

type OrderRepositoryDB struct {
	client *pgx.Conn
}

func (s OrderRepositoryDB) GetAll(tenantId int, userName string) ([]domain.Order, error) {

	rows, err := s.client.Query(context.Background(),
		`select o.order_id, p."name", op.product_quantity, p.price
	from "order" o 
	inner join order_product op on o.order_id = op.order_id 
	inner join product p on op.product_id = p.product_id 
	inner join users u on o.user_name = u.username 
	where u.tenant_id=$1 and u.username = $2`, tenantId, userName)
	if err != nil {
		fmt.Printf("Query failed: %v\n", err)
		return nil, fmt.Errorf("Query failed: %v\n", err)
	}

	//fmt.Printf("Query count: %v\n", rows.))
	orders := []domain.Order{}
	for rows.Next() {
		var order_id int
		var product_name string
		var product_quantity int
		var price int
		err = rows.Scan(&order_id, &product_name, &product_quantity, &price)
		if err != nil {
			fmt.Printf("Query scan failed: %v\n", err)
			return nil, fmt.Errorf("Query scan failed: %v\n", err)
		}
		d := domain.Order{
			OrderId: order_id,
			Products: []domain.Product{
				{
					ProductName: product_name,
					Price:       float64(price),
					Quantity:    product_quantity,
				},
			},
		}
		orders = append(orders, d)
	}

	return orders, nil
}

func NewOrderRepositoryDB() OrderRepositoryDB {

	var err error
	if databaseconn == nil {
		// urlExample := "postgres://postgres:password@localhost:5432/postgres"
		databaseconn, err = pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
		//defer conn.Close(context.Background())
	}

	return OrderRepositoryDB{client: databaseconn}
}
