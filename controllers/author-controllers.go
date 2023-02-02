package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dhruv-vavliya/BookStore/models"
	"github.com/dhruv-vavliya/BookStore/services"
	"github.com/golang-jwt/jwt/v4"
)

func SignUp(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	params := models.Author{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		json.NewEncoder(w).Encode( models.ResponseError{
			Status: http.StatusBadRequest,
			Message: err.Error(),
		} )
		return
	}
	fmt.Print(params)

	user, err := services.RegisterAuthor(r.Context(), EntClient, &params )
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(user)
}


func Login(w http.ResponseWriter, r *http.Request){

	// validate username & password.
	params := &models.Credentials{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		json.NewEncoder(w).Encode( models.ResponseError{
			Status: http.StatusUnauthorized,
			Message: err.Error(),
		} )
		return
	}

	authorID, err := services.ValidateAuthor(r.Context(), EntClient, params)
	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusUnauthorized,
			Message: err.Error(),
		})	
		return	
	}

	// sign JWT
	secret_key := []byte(os.Getenv("JWT_SECRET_KEY"))
	expirationTime := time.Now().Add(time.Hour * 24 * 30)

	// create JWT token
	claims := &models.Claims{
		UserID: authorID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),	
		},
	}

	// sign new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signature, err := token.SignedString(secret_key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	// store token to client's cookie storage.
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: signature,
		Expires: expirationTime,
	})
}


func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Expires: time.Now(),
	})
}


func DeleteAccount(w http.ResponseWriter, r *http.Request){
	authorID, _ := strconv.Atoi(r.Header.Get("id"))
	err := services.DeleteAuthorByID(r.Context(), EntClient, authorID)

	if err != nil {
		json.NewEncoder(w).Encode(models.ResponseError{
			Status: http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
}

