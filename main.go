package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"to-do-backend/database"
	"to-do-backend/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Get server port from environment
	serverPort := os.Getenv("SERVER_PORT")
	serverHost := os.Getenv("SERVER_HOST")
	if serverPort == "" {
		log.Fatal("SERVER_PORT not set in environment")
	}

	// Initialize database with environment variables
	dbConfig := database.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db := database.InitDB(dbConfig)

	logger := log.New(os.Stdout, "[TODO-APP] ", log.LstdFlags)

	// Initialize the server
	srv := server.NewServer(serverHost +":"+serverPort, db, logger)

	// Run the server in a goroutine
	go func() {
		log.Println("Starting server on port", serverHost +":"+serverPort)
		if err := srv.Start(); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a timeout context for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.HTTPServer.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
