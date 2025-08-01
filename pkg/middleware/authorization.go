package middleware

import (
	"context"
	"go/http/configs"
	"go/http/pkg/jwt"
	"log"
	"net/http"
	"strings"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func WriteUnauthed(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")

		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			WriteUnauthed(w)
			return
		}
		token := strings.TrimPrefix(authorizationHeader, "Bearer ")
		log.Println(token)
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)

		if !isValid {
			WriteUnauthed(w)
			return
		}
		ctx := context.WithValue(r.Context(), ContextEmailKey, data.Email)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
