package controllers

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	SAPID      string    `json:"sap_id"`
	UserID     int       `json:"user_id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	Department string    `json:"department"`
	Expires    time.Time `json:"expires"`
	jwt.RegisteredClaims
}
