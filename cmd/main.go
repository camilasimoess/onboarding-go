package main

import (
	"github.com/camilasimoess/onboarding-go/db"
	"github.com/camilasimoess/onboarding-go/internal/handler"
	"github.com/camilasimoess/onboarding-go/internal/repo"
	"github.com/camilasimoess/onboarding-go/internal/service"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	database := db.NewDatabase()
	userRepository := repo.NewUserRepository(database, "users")
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("POST /save", userHandler.CreateUser)
	http.HandleFunc("GET /find/{id}", userHandler.GetUser)

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
