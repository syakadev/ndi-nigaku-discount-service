package http_util

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM) // Tangkap lebih banyak sinyal

		<-sigint
		log.Println("Shutdown signal received...")

		// Graceful shutdown dengan timeout 5 detik
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := a.ShutdownWithContext(ctx); err != nil {
			log.Printf("Error shutting down server: %v", err)
		} else {
			log.Println("Server stopped gracefully.")
		}

		close(idleConnsClosed)
	}()

	// Jalankan server
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Fatalf("Server error: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
