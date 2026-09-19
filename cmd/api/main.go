package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/fahreyad/go_api/internal/config"
	"github.com/fahreyad/go_api/internal/db"
	"github.com/fahreyad/go_api/internal/handlers"
	"github.com/fahreyad/go_api/internal/middleware"
)

func main() {
	cnf := config.MustLoad()
	db, err := db.Connect(cnf.DATABASE_URL)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// logger setup
	loghandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo, AddSource: true})
	logger := slog.New(loghandler)
	slog.SetDefault(logger)

	//router set up
	mux := http.NewServeMux()

	handler := handlers.NewHandler(db, logger)
	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("GET /listings", handler.List)
	mux.HandleFunc("DELETE /listings/{id}", handler.Delete)

	requestIDMiddleware := middleware.RequestID(mux)
	// server set up
	serv := &http.Server{
		Addr:         ":" + cnf.Port,
		Handler:      requestIDMiddleware,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Starting server on %s", serv.Addr)

	if err := serv.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
