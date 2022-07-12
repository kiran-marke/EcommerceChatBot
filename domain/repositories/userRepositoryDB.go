package domain

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	domain "github.com/kiran-marke/ecommercechatbot/domain/models"
)

type UserRepositoryDB struct {
	client *pgx.Conn
	user   domain.User
}

func (u UserRepositoryDB) GetUserDetails() (domain.User, error) {

	user := domain.User{
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
