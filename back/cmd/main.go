package main

import (
	"Flowers-store/internal/handler"
	"Flowers-store/internal/repository"
	"Flowers-store/internal/service"
	"Flowers-store/pkg/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.NewDB()

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close()

	repo := repository.NewPostgresRepository(db)
	authSvc := service.NewAuthService(repo)
	regSvc := service.NewRegistrationService(repo)

	router := mux.NewRouter()

	router.HandleFunc("/auth", handler.NewAuthHandler(authSvc).HandleLogin).Methods("POST")
	router.HandleFunc("/register", handler.NewRegistrationHandler(regSvc).HandleRegister).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server is running at localhost:8080")
	log.Fatal(server.ListenAndServe())
}
