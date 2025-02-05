package main

import (
	"net/http"
	"onboarding-go/internal/handler"
	"onboarding-go/internal/repo"
	"onboarding-go/internal/service"
)

func main() {
	userRepository := repo.NewUserRepository()
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/save", userHandler.SaveUser)
	http.HandleFunc("/find", userHandler.FindUserByID)
}
