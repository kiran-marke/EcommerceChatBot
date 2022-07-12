package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	domain "github.com/kiran-marke/ecommercechatbot/domain/models"
	"github.com/kiran-marke/ecommercechatbot/service"
)

type Handler struct {
	service service.Service
}

func (h *Handler) Greeting(w http.ResponseWriter, r *http.Request) {

	claims, err := h.checkForAuthorizationToken(r)
	if err != nil {
		writeResponse(w, http.StatusForbidden, fmt.Sprintf("error in fetching greeting: %v", err))
		return
	}

	var greetingMessage string
	greetingMessage = fmt.Sprintf(`Welcome %v to the %v Super Shop. 
	Do you want to perform one of the following options? 
	1.Aboutus 2.ShowOrders`, claims.Username, claims.Tenantname)
	w.Header().Add("Content-Type", "application-json")
	json.NewEncoder(w).Encode(domain.Greeting{GreetingMessage: greetingMessage})
}

func (h *Handler) checkForAuthorizationToken(r *http.Request) (*domain.AccessTokenClaims, error) {
	token := r.Header.Get("Authorization")
	if token != "" {
		d := domain.VerifyRequest{
			AuthToken: token,
		}
		claims, err := h.service.Verify(d)

		if err != nil {
			return nil, fmt.Errorf("missing token")
		}

		return claims, nil
	}

	return nil, fmt.Errorf("missing token")
}

func (h *Handler) PerformTask(w http.ResponseWriter, r *http.Request) {

	claims, err := h.checkForAuthorizationToken(r)
	if err != nil {
		writeResponse(w, http.StatusForbidden, fmt.Sprintf("error in performing task: %v", err))
		return
	}

	taskName := r.Header.Get("TaskName")
	if taskName == "" {
		writeResponse(w, http.StatusBadRequest, "missing headers")
		return
	}

	switch strings.ToLower(taskName) {
	case "aboutus":
		tenant, _ := h.service.GetTenantDetails(claims.TenantId)
		w.Header().Add("Content-Type", "application-json")
		json.NewEncoder(w).Encode(tenant)

	case "showorders":
		orders, _ := h.service.GetAllOrders(claims.TenantId, claims.Username)
		w.Header().Add("Content-Type", "application-json")
		json.NewEncoder(w).Encode(orders)

	default:
		w.Header().Add("Content-Type", "application-json")
		json.NewEncoder(w).Encode("Please select available menu only.")
	}

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login request arrived!")

	var loginRequest domain.LoginRequest
	username := r.Header.Get("username")
	password := r.Header.Get("password")

	if len(username) <= 0 || len(password) <= 0 {
		writeResponse(w, http.StatusBadRequest, errors.New("Invalid credentials"))
		return
	}

	loginRequest = domain.LoginRequest{
		Username: username,
		Password: password,
	}

	token, appErr := h.service.Login(loginRequest)
	if appErr != nil {
		writeResponse(w, http.StatusUnauthorized, appErr.Error())
	} else {
		writeResponse(w, http.StatusOK, *token)
	}
}

func (h *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("authorization")

	verifyRequest := domain.VerifyRequest{
		AuthToken: authToken,
	}

	if authToken != "" {
		claims, appErr := h.service.Verify(verifyRequest)
		if appErr != nil {
			writeResponse(w, http.StatusUnauthorized, appErr.Error())
		} else {
			writeResponse(w, http.StatusOK, claims)
		}
	} else {
		writeResponse(w, http.StatusForbidden, "missing token")
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
