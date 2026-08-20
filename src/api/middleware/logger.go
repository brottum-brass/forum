package middleware

import (
	"log"
	"net/http"
)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Next(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("INCOMMING REQUEST FROM %s:\t%s\t%s\n", r.RemoteAddr, r.Method, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
