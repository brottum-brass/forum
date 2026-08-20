package middleware

import "net/http"

type Middleware interface {
	Next(next http.Handler) http.Handler
}
