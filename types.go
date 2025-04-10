package types

import "github.com/dgrijalva/jwt-go"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
