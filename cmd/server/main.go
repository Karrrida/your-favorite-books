package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/Karrrida/your-favorite-books/internal/config"
	"github.com/Karrrida/your-favorite-books/internal/database"
	"github.com/Karrrida/your-favorite-books/internal/handler"
	"github.com/Karrrida/your-favorite-books/internal/repository"
	"github.com/Karrrida/your-favorite-books/internal/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}
	
	db, err := database.NewPostgres(cfg.DatabaseURL)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connection success")
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handler.Handler)
	r.Get("/info", handler.GetServerOsInfo)
	r.Post("/register", userHandler.Register)
	r.Get("/get-by-email", userHandler.GetByEmail)

	
	address := ":" + cfg.Port
	fmt.Printf("Starting server on %s\n", address)

	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatal(err)
	}
}
