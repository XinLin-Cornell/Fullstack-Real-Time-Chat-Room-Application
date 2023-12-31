package middlewares

import (
	"net/http"
	"os"

	"github.com/rs/cors"
)

func Cors(loggedRoutes http.Handler) http.Handler {
	allowedHeaders := []string{"Authorization", "Content-Type"}
	allowedOrigins := []string{"http://localhost:3000"} // specify the domain of your frontend here
	allowedMethods := []string{"GET", "POST", "PUT", "DELETE"}

	corsDebug := os.Getenv("CORS_DEBUG")
	if corsDebug == "true" {
		return cors.New(cors.Options{
			Debug:          true,
			AllowedHeaders: allowedHeaders,
			AllowedOrigins: allowedOrigins,
			AllowedMethods: allowedMethods,
		}).Handler(loggedRoutes)
	}
	return cors.New(cors.Options{
		AllowedHeaders: allowedHeaders,
		AllowedOrigins: allowedOrigins,
		AllowedMethods: allowedMethods,
	}).Handler(loggedRoutes)
}
