package routes

import (
	"github.com/dhruv-vavliya/BookStore/controllers"
	"github.com/dhruv-vavliya/BookStore/middlewares"
	"github.com/gorilla/mux"
)

func RegisterBookRoutes(router *mux.Router) {
	router.HandleFunc("/book", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{author}", controllers.GetBooksByAuthorName).Methods("GET")

	bookRouter := router.PathPrefix("/book").Subrouter()
	bookRouter.Use(middlewares.Auth)
	bookRouter.HandleFunc("/publish", controllers.AddBook).Methods("POST")
	bookRouter.HandleFunc("/delete/{id}", controllers.DeleteBook).Methods("DELETE")
	bookRouter.HandleFunc("/update/{id}", controllers.UpdateBook).Methods("PUT")
}
