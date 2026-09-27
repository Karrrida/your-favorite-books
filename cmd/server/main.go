package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Karrrida/your-favorite-books/internal/database"
	"github.com/Karrrida/your-favorite-books/internal/handler"
	"github.com/Karrrida/your-favorite-books/internal/repository"
	"github.com/Karrrida/your-favorite-books/internal/service"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env")
	}

	databseUrl := os.Getenv("DATABASE_URL")

	db, err := database.NewPostgres(databseUrl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db connection success")
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/info", handler.GetServerOsInfo)
	http.HandleFunc("POST /register", userHandler.Register)
	http.HandleFunc("/get-by-email", userHandler.GetByEmail)

	fmt.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
