package middleware

import (
	"github.com/rs/cors"
	"net/http"
)

func Cors() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"}, // разрешает все домены
			AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		})
		return c.Handler(next)
	}
}
