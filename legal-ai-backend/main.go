package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yourusername/legal-ai-backend/handlers"
	"github.com/yourusername/legal-ai-backend/utils"
)

func main() {
	utils.InitLogger()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/upload", handlers.UploadHandler)

	log.Printf("starting server on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
