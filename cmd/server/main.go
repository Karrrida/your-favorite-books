package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Karrrida/your-favorite-books/internal/database"
	"github.com/Karrrida/your-favorite-books/internal/handler"
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

	http.HandleFunc("/", handler.Handler)
	http.HandleFunc("/info", handler.GetServerOsInfo)
	http.HandleFunc("/register", handler.Register)

	fmt.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
