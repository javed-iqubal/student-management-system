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
	"github.com/javed-iqubal/student-management-system/internal/http/handler/student"
	"github.com/javed-iqubal/student-management-system/internal/storage/sqlite"
)

func main() {

	// load config

	cfg := config.MustLoad()

	//Database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage initialized ", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// Setup Router

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))
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
			log.Fatal("Failed to start server", err)
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
