package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/dhruv-vavliya/BookStore/models"
	"github.com/dhruv-vavliya/BookStore/services"
	"github.com/gorilla/mux"
)

func AddBook(w http.ResponseWriter, r *http.Request) {

	params := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})	
		return	
	}

	authorID, _ := strconv.Atoi(r.Header.Get("id"))
	author, _ := EntClient.Author.Get(r.Context(), authorID)

	book, err := services.CreateBook(r.Context(), EntClient, params, author)
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(book)
}


func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, _ := strconv.Atoi(params["id"])
	
	ctx := r.Context()
	ctx = context.WithValue(ctx, "bookID", bookID)

	err := services.DeleteBookByID(ctx, EntClient)
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})	
		return	
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	bookID, _ := strconv.Atoi(mux.Vars(r)["id"])

	params := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})	
		return	
	}

	err = services.UpdateBookByID(r.Context(), EntClient, bookID, params)
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})	
		return	
	}

	w.WriteHeader(http.StatusOK)
}


func GetAllBooks(w http.ResponseWriter, r *http.Request) {

	books, err := services.GetAllBooks(r.Context(), EntClient)
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})	
		return	
	}

	json.NewEncoder(w).Encode(books)
}

func GetBooksByAuthorName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	authorName := params["author"]

	books, err := services.GetBooksByAuthorName(r.Context(), EntClient, authorName)
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		})	
		return	
	}

	json.NewEncoder(w).Encode(books)
}

