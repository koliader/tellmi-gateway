package token

type Maker interface {
	VerifyToken(tokenString string) (*Payload, error)
}
