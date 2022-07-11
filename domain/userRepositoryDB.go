package domain

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type UserRepositoryDB struct {
	client *pgx.Conn
	user   User
}

func (u UserRepositoryDB) GetUserDetails() (User, error) {

	// var tenant_id int
	// var tenant_name string
	// var tenant_details string
	// err := t.client.QueryRow(context.Background(), "select tenant_id, tenant_name, tenant_details from user where tenant_id=$1", 1).Scan(&tenant_id, &tenant_name, &tenant_details)
	// if err != nil {
	// 	log.Fatalf("QueryRow failed: %v\n", err)
	// }

	user := User{
		TenantId: 1,
		UserName: "amazonuser",
	}
	u.user = user

	return u.user, nil
}

func NewUserRepositoryDB() UserRepositoryDB {

	var err error
	if databaseconn == nil {
		databaseconn, err = pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
		//defer conn.Close(context.Background())
	}

	return UserRepositoryDB{client: databaseconn}
}
