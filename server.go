package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dhruv-vavliya/BookStore/configs"
	"github.com/dhruv-vavliya/BookStore/controllers"
	"github.com/dhruv-vavliya/BookStore/middlewares"
	"github.com/dhruv-vavliya/BookStore/routes"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Bookstore</h1>"))
}

// Router -> routes -> controllers -> services -> repos -> modelling -> database
func main() {

	// Database connection
	client, err := configs.Connect()
	if err != nil {
		log.Fatal("Database Connection Failed.")
	}
	controllers.EntClient = client
	middlewares.EntClient = client

	log.Println("Database Connected Successfully.")
	defer client.Close()

	// Router Setup
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	routes.RegisterBookRoutes(r)
	routes.RegisterAuthorRoutes(r)

	// Server Setup
	srv := &http.Server{
		Handler:      r,
		Addr:         os.Getenv("HOST_ADDR"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server Port: %s", os.Getenv("PORT"))
	log.Fatal(srv.ListenAndServe())
}
