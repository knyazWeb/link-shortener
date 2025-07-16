package middleware

import (
	"log"
	"net/http"
	"strings"
)

func BearerToken(next http.Handler) http.Handler {
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
		next.ServeHTTP(w, r)
	})
}
