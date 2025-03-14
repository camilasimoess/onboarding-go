package main

import (
	"log/slog"
	"net/http"
	"onboarding-go/db"
	"onboarding-go/internal/handler"
	"onboarding-go/internal/repo"
	"onboarding-go/internal/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	database := db.NewDatabase()
	userRepository := repo.NewUserRepository(database, "onboarding", "users")
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("POST /save", userHandler.SaveUser)
	http.HandleFunc("GET /find/{id}", userHandler.FindUserByID)

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			slog.Error("error starting server: ", err)
			return
		}
	}()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	slog.Info("server stopped")
}
