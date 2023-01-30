package routes

import (
	"github.com/dhruv-vavliya/BookStore/controllers"
	"github.com/dhruv-vavliya/BookStore/middlewares"
	"github.com/gorilla/mux"
)

func RegisterAuthorRoutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.SignUp ).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	authorRouter := router.PathPrefix("/author").Subrouter()
	authorRouter.Use(middlewares.Auth)
	authorRouter.HandleFunc("/delete", controllers.DeleteAccount).Methods("DELETE")
}
