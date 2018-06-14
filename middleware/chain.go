package middleware

import (
	"net/http"
)

// Middleware defines the http middleware type.
type Middleware func(http.Handler) http.Handler

// Chain defines a middleware chain.
type Chain struct {
	middlewares []Middleware
}

// NewChain returns a new middleware chain.
func NewChain(mws ...Middleware) *Chain {
	c := &Chain{make([]Middleware, len(mws))}
	for i, mw := range mws {
		c.middlewares[i] = mw
	}
	return c
}

// Add adds middlewares to the middleware chain.
func (c *Chain) Add(mws ...Middleware) {
	if c.middlewares == nil {
		c.middlewares = make([]Middleware, 0)
	}
	c.middlewares = append(c.middlewares, mws...)
}

// Final does the actual chaining of all the middlewares present with it
// and returns the final http.Handler that can be used to handle a request.
func (c *Chain) Final(h http.Handler) http.Handler {
	final := h
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		final = c.middlewares[i](final)
	}
	return final
}
