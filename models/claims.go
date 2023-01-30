package models

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserID int `json:"userid"`
	// Pre-Defined Claims
	jwt.RegisteredClaims
}