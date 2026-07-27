package middleware

import "github.com/koliader/tellmi-sdk/token"

type Middleware struct {
	tokenMaker token.Maker
}

func NewMiddleware(tokenMaker token.Maker) *Middleware {
	return &Middleware{
		tokenMaker: tokenMaker,
	}
}
