package middleware

import (
	"go/http/configs"
	"go/http/pkg/jwt"
	"log"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		splitedHeader := strings.Split(authorizationHeader, " ")

		if len(splitedHeader) < 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := splitedHeader[1]
		log.Println(token)
		isValid, data := jwt.NewJWT(config.Auth.Secret).Parse(token)
		log.Println(isValid)
		log.Println(data)
		next.ServeHTTP(w, r)
	})
}
