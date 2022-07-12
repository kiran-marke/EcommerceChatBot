package service

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	domain "github.com/kiran-marke/ecommercechatbot/domain/models"
)

type Service interface {
	GetAllOrders(int, string) ([]domain.Order, error)
	GetUserDetails() (domain.User, error)
	GetTenantDetails(int) (domain.Tenant, error)
	Login(domain.LoginRequest) (*domain.LoginResponse, error)
	Verify(domain.VerifyRequest) (*domain.AccessTokenClaims, error)
}

type DefaultService struct {
	orderRepo  domain.OrderRepository
	userRepo   domain.UserRepository
	tenantRepo domain.TenantRepository
	authrepo   domain.AuthRepository
}

func (s DefaultService) GetAllOrders(tenantId int, userName string) ([]domain.Order, error) {
	return s.orderRepo.GetAll(tenantId, userName)
}

func (s DefaultService) GetUserDetails() (domain.User, error) {
	return s.userRepo.GetUserDetails()
}

func (s DefaultService) GetTenantDetails(tenantId int) (domain.Tenant, error) {
	return s.tenantRepo.GetTenantDetails(tenantId)
}

func (s DefaultService) Login(req domain.LoginRequest) (*domain.LoginResponse, error) {
	var appErr error
	var login *domain.Login

	if login, appErr = s.authrepo.FindBy(req.Username, req.Password); appErr != nil {
		return nil, appErr
	}

	claims := login.ClaimsForAccessToken()
	authToken := domain.NewAuthToken(claims)

	var accessToken string
	if accessToken, appErr = authToken.NewAccessToken(); appErr != nil {
		return nil, appErr
	}

	return &domain.LoginResponse{AccessToken: accessToken}, nil
}

func (s DefaultService) Verify(req domain.VerifyRequest) (*domain.AccessTokenClaims, error) {

	// convert the string token to JWT struct
	if jwtToken, claims, err := jwtTokenFromString(req.AuthToken); err != nil {
		return nil, err
	} else {
		if jwtToken.Valid {

			return claims, nil
		} else {
			return nil, fmt.Errorf("Invalid token")
		}
	}
}

func jwtTokenFromString(tokenString string) (*jwt.Token, *domain.AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})

	if err != nil {
		log.Println("Error while parsing token: " + err.Error())
		return nil, nil, err
	}
	claims := token.Claims.(*domain.AccessTokenClaims)
	return token, claims, nil
}

func NewService(or domain.OrderRepository, ur domain.UserRepository, tr domain.TenantRepository, ar domain.AuthRepository) DefaultService {
	return DefaultService{
		orderRepo:  or,
		userRepo:   ur,
		tenantRepo: tr,
		authrepo:   ar,
	}
}
