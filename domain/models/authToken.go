package domain

import (
	"fmt"

	"log"

	"github.com/dgrijalva/jwt-go"
)

type AuthToken struct {
	token *jwt.Token
}

type AuthRepository interface {
	FindBy(username string, password string) (*Login, error)
}

func (t AuthToken) NewAccessToken() (string, error) {
	signedString, err := t.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		log.Println("Failed while signing access token: " + err.Error())
		return "", fmt.Errorf("cannot generate access token")
	}
	return signedString, nil
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}
