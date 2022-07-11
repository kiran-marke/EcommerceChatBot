package domain

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type AuthRepository interface {
	FindBy(username string, password string) (*Login, error)
}

type AuthRepositoryDb struct {
	client *pgx.Conn
}

func (d AuthRepositoryDb) FindBy(username, password string) (*Login, error) {
	var login Login
	var user_name string
	var tenant_id int
	var tenant_name string

	sqlVerify := `SELECT username, u.tenant_id, t.tenant_name 
					FROM users u
					inner join tenant t on u.tenant_id = t.tenant_id 
					WHERE username = $1 and password = $2`
	err := d.client.QueryRow(context.Background(), sqlVerify, username, password).Scan(&user_name, &tenant_id, &tenant_name)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	login = Login{
		TenantId:   tenant_id,
		Username:   user_name,
		TenantName: tenant_name,
	}
	return &login, nil
}

func NewAuthRepository() AuthRepositoryDb {
	var err error
	if databaseconn == nil {
		databaseconn, err = pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
		//defer conn.Close(context.Background())
	}
	return AuthRepositoryDb{client: databaseconn}
}
