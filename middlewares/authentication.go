package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/dhruv-vavliya/BookStore/ent"
	"github.com/dhruv-vavliya/BookStore/models"
	"github.com/golang-jwt/jwt/v4"
)

var EntClient *ent.Client

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request) {

		tokenstring, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		signature := tokenstring.Value
		secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		claims := &models.Claims{}
		
		// parse tokenstring & store to claims
		token, err := jwt.ParseWithClaims(signature, claims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		
		// if token has expired or invalid
		if err != nil {
			if err == jwt.ErrTokenExpired {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// if token isn't valid
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// redirect to original route
		r.Header.Set("id", strconv.Itoa(claims.UserID))
		fmt.Println("Authentication Done.")
		next.ServeHTTP(w, r)
	})
}