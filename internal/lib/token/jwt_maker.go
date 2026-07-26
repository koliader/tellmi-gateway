package token

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

const minSecretKeySize = 32

var ErrInvalidToken = errors.New("invalid token")

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %v", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (m *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Payload{},
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}
			return []byte(m.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Payload)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
