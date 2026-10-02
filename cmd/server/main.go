package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fizz-buzz/internal/db"
	"fizz-buzz/internal/handler"
)

const Version = "v0.1.0"
const Author = "Najeh Abassi"

func setupRouter(database *db.DB) http.Handler  {
	h := handler.NewHTTPHandler(database)
	mux := http.NewServeMux()

	// Used when calling with curl but passing params as string exemple ?int1=3&int2=6..
	mux.HandleFunc("GET /api/v1/fizzbuzz", h.HandleFizzBuzzQuery)
	mux.HandleFunc("GET /api/v1/stats", h.HandleStats)

	return mux
}

func main() {
	// Initialisation de base de donnée pour utilisation des rapports et statistique
	database, err := db.InitDB("fizzbuzz.db")
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer database.Close()

	// Used when calling with curl post and providing json params for testing only to be removed
	//mux.HandleFunc("POST /api/v1/fizzbuzz", h.HandleFizzBuzz)

	server := &http.Server{
		Addr:         ":8081",
		Handler:      setupRouter(database),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		slog.Info("Server listening", "port", 8081)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server error", "error", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("Shutting down server gracefully...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced shutdown", "error", err)
	}


}