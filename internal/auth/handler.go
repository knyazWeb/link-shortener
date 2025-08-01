package auth

import (
	"go/http/configs"
	"go/http/pkg/jwt"
	"go/http/pkg/request"
	"go/http/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := request.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}

		email, err := handler.AuthService.Login(body.Email, body.Password)

		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		jwtString, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})

		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		data := LoginResponse{
			Token: jwtString,
		}

		response.Json(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}

		email, err := handler.AuthService.Register(body.Name, body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		jwtString, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Email: email})

		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		data := RegisterResponse{
			Token: jwtString,
		}

		response.Json(w, data, http.StatusOK)
	}
}
