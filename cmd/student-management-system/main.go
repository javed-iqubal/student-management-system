package main

import (
	"fmt"
	"log"
	"net/http"

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
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}
