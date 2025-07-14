package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "github.com/lo0kat/kerebro-api/db"
    "github.com/lo0kat/kerebro-api/router"
)

func main() {
    db.Connect(os.Getenv("DATABASE_URL"))
    defer db.Pool.Close() // Close when main exits

	// Graceful shutdown on SIGINT/SIGTERM
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

    server := &http.Server{
        Addr:    ":8080",
        Handler: router.NewRouter(),
    }   
    go func() {
        log.Printf("Server running on port :8080")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Server failed: %v", err)
        }
    }()

    <-quit
    log.Println("Shutting down server...")

    if err := server.Close(); err != nil {
        log.Fatalf("Server shutdown failed: %v", err)
    }

    log.Println("Server gracefully stopped")
}
