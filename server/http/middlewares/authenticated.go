package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"chat/http/responses"
	"chat/utils"
	"chat/utils/errors"

	"github.com/golang-jwt/jwt"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		jwtSecret := os.Getenv("JWT_SECRET")
		if len(authHeader) != 2 {
			handleAuthenticationErr(w, errors.ErrMalformedToken)
			return
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(jwtSecret), nil
			})

			if err != nil {
				handleAuthenticationErr(w, err)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				var props utils.JWTProps = "JWTProps"
				ctx := context.WithValue(r.Context(), props, claims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				handleAuthenticationErr(w, err)
				return
			}
		}
	})
}

func handleAuthenticationErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := responses.ErrorResponse{Message: err.Error(), Status: false, Code: http.StatusUnauthorized}
	data, err := json.Marshal(res)
	errors.ErrorCheck(err)
	w.Write(data)
}
