package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type Payload struct {
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return jwt.NewValidationError("token is expired", jwt.ValidationErrorExpired)
	}
	return nil
}
