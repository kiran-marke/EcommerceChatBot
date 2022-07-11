package domain

type Tenant struct {
	TenantId   int    `json:"tenant_id"`
	TenantName string `json:"tenant_name"`
	AboutUs    string `json:"aboutus"`
}

type TenantRepository interface {
	GetTenantDetails(int) (Tenant, error)
}
