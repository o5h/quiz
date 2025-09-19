package token

import "github.com/golang-jwt/jwt/v5"

// Claims represents the custom claims for JWT tokens.
type Claims struct {
	UserName string `json:"user_name"`
	Role     string `json:"role"`
}

// AccessClaims represents the JWT claims for access tokens.
type AccessClaims struct {
	Claims
	jwt.RegisteredClaims
}
