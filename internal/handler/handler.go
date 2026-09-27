package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/Karrrida/your-favorite-books/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

type RegisterBody struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world, you've requested: %s\n", r.URL.Path)
}

func GetServerOsInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server use number of cpus: %d\n", runtime.NumCPU())
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterBody
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid params", http.StatusBadRequest)
		return
	}

	err = h.userService.Register(req.Email, req.Name)
	if err != nil {
		fmt.Println("ERROR IN HANDER", err)
		http.Error(w, "Invalid params", http.StatusBadRequest)
	}
}

func (h *UserHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	user, err := h.userService.GetByEmail(email)

	if err != nil {
		fmt.Println("Error in UserHandler GetByEmail", err)
		http.Error(w, "User not found", http.StatusNotFound)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("Error occured in UserHandler GetByEmail", err)
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
}
