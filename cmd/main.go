package main

import (
	"context"
	"github.com/camilasimoess/onboarding-go/internal/db"
	"github.com/camilasimoess/onboarding-go/internal/handler"
	"github.com/camilasimoess/onboarding-go/internal/repo"
	"github.com/camilasimoess/onboarding-go/internal/service"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	client, database := db.NewDatabase()
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := client.Disconnect(ctx); err != nil {
			slog.Error("error disconnecting from MongoDB:", err)
		} else {
			slog.Info("disconnected from MongoDB successfully")
		}
	}()

	userRepository := repo.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		userHandler.CreateUser(context.Background(), w, r)
	})
	http.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		userHandler.GetUser(context.Background(), w, r)
	})

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
