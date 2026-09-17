package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fahreyad/go_api/internal/config"
	"github.com/fahreyad/go_api/internal/db"
	"github.com/fahreyad/go_api/internal/handlers"
)

func main() {
	cnf := config.MustLoad()
	_, err := db.Connect(cnf.DATABASE_URL)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Successfully connected to database")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)
	serv := &http.Server{
		Addr:         ":" + cnf.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Starting server on %s", serv.Addr)

	if err := serv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
