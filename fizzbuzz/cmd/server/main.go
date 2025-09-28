package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/fizzbuzz/internal/config"
	"github.com/fizzbuzz/internal/handlers"
	"github.com/fizzbuzz/internal/repository"
	"github.com/fizzbuzz/internal/services"
)

func main() {
	ctx, cancelFn := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	defer cancelFn()

	if err := run(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v", err)
		cancelFn()
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg := config.Load()

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	repository := repository.NewRepository(db)

	service := services.NewService(repository)

	r := handlers.SetupRouter(service)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		log.Printf("Server starting on port %s", cfg.Port)
		serverErr <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErr:
		if err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("server failed to start: %w", err)
		}
	case <-ctx.Done():
		log.Println("Shutting down server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server forced to shutdown: %w", err)
		}

		log.Println("Server exited")
	}

	return nil
}

func connectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPass + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"
	return sql.Open("postgres", dsn)
}
