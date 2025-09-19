package token

import "errors"

var ErrInvalidToken = errors.New("invalid token")
var ErrTokenInvalidClaims = errors.New("token has invalid claims")
