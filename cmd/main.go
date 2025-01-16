package main

import (
	"mine/internal/handler"
	"mine/internal/repo"
	"mine/internal/service"
	"net/http"
)

func main() {
	userRepository := repo.NewUserRepository()
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/save", userHandler.SaveUser)
	http.HandleFunc("/find", userHandler.FindUserByID)
}
