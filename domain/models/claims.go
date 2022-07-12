package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"
const ACCESS_TOKEN_DURATION = time.Hour
const REFRESH_TOKEN_DURATION = time.Hour * 24 * 30

type AccessTokenClaims struct {
	TenantId   int    `json:"tenant_id"`
	Tenantname string `json:"tenant_name"`
	Username   string `json:"username"`
	jwt.StandardClaims
}

func (c AccessTokenClaims) IsValidTenantId(tenantId int) bool {
	return c.TenantId == tenantId
}
