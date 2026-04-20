package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/javed-iqubal/student-management-system/internal/config"
)

func main() {

	// load config

	cfg := config.MustLoad()

	// Database setup

	// Setup Router

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to student management system!"))
	})
	// Setup Server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Println("Server started")

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()

	<-done

	slog.Info("Server is shutting down")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server failed to shutdown ", slog.String("error", err.Error()))
	}

	slog.Info("Server shutdown successfully!")
}
