package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	domain "github.com/kiran-marke/ecommercechatbot/domain/repositories"
	"github.com/kiran-marke/ecommercechatbot/service"
	"github.com/rs/cors"
)

func Start() {

	router := mux.NewRouter()
	//wiring
	h := Handler{
		service: service.NewService(
			domain.NewOrderRepositoryDB(),
			domain.NewUserRepositoryDB(),
			domain.NewTenantRepositoryDB(),
			domain.NewAuthRepository()),
	}

	//define routes
	router.HandleFunc("/greeting", h.Greeting).Methods(http.MethodGet)
	router.HandleFunc("/performtask", h.PerformTask).Methods(http.MethodPost)

	router.HandleFunc("/auth/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", h.Verify).Methods(http.MethodGet)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"username", "password", "Content-Type", "Authorization", "TaskName"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS"},
	})

	handler := c.Handler(router)

	//starting the server
	log.Fatal(http.ListenAndServe("localhost:8000", handler))
}
