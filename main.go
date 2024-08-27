package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5" // Import the chi package
	"github.com/go-chi/cors"   // Import the cors package
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("chandni")

	if err := godotenv.Load(".env"); err != nil {
		log.Printf("No .env file found, using default environment variables")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found")
	}

	router := chi.NewRouter()

	// Use CORS middleware
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	srv := &http.Server{Handler: router, Addr: ":" + portString}

	// Define a simple route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

v1Router:=chi.NewRouter()
v1Router.Get("/healthz",handlerReadiness)
v1Router.Get("/readiness",handlerReadiness)
v1Router.Get("/err",handlerErr)

router.Mount("/v1",v1Router)

	log.Printf("Server running at %v", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
