package domain

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
	domain "github.com/kiran-marke/ecommercechatbot/domain/models"
)

type TenantRepositoryDB struct {
	client *pgx.Conn
}

func (t TenantRepositoryDB) GetTenantDetails(tenantId int) (domain.Tenant, error) {

	var tenantName string
	var tenantDetails string
	err := t.client.QueryRow(context.Background(), "select tenant_name, tenant_details from tenant where tenant_id=$1", tenantId).Scan(&tenantName, &tenantDetails)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	tenantStruct := domain.Tenant{
		TenantId:   tenantId,
		TenantName: tenantName,
		AboutUs:    tenantDetails,
	}

	return tenantStruct, nil
}

func NewTenantRepositoryDB() TenantRepositoryDB {
	var err error
	if databaseconn == nil {
		// urlExample := "postgres://postgres:password@localhost:5432/postgres"
		databaseconn, err = pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
		if err != nil {
			log.Fatalf("Unable to connect to database: %v\n", err)
		}
		//defer conn.Close(context.Background())
	}

	return TenantRepositoryDB{client: databaseconn}
}
