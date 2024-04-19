package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT environment variable was not set")
	}

	fmt.Println("PORT is set to", portString)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
	}))

	v1Router := chi.NewRouter()

	router.Mount("/v1", v1Router)

	v1Router.Get("/health", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server is running on port %s", portString)

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
