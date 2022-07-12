package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Login struct {
	Username   string `db:"username"`
	TenantId   int    `db:"tenant_id"`
	TenantName string `db:"tenant_name"`
}

func (l Login) ClaimsForAccessToken() AccessTokenClaims {

	return l.claimsForUser()
}

func (l Login) claimsForUser() AccessTokenClaims {
	return AccessTokenClaims{
		TenantId:   l.TenantId,
		Username:   l.Username,
		Tenantname: l.TenantName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}
