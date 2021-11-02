package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	ID        uint
	FirstName string
	LastName  string
	Email       string
	Mobile string
	jwt.StandardClaims
}