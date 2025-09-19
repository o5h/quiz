package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAccessToken(user Claims, duration time.Duration) (string, error) {
	claims := AccessClaims{
		Claims: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "quiz_app",
			Subject:   user.UserName,
			ID:        "",
			Audience:  nil,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(JWT_KEY)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ParseAccessToken(tokenStr string) (*AccessClaims, error) {
	claims := &AccessClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})
	if err != nil {
		if err == jwt.ErrTokenInvalidClaims {
			return nil, ErrTokenInvalidClaims
		}
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
