package auth

import (
	"go/http/configs"
	"go/http/pkg/request"
	"go/http/pkg/response"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, err := request.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}

		res := LoginResponse{
			Token: "123",
		}

		response.Json(w, res, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}

		res := RegisterResponse{
			Token: "123",
		}

		response.Json(w, res, http.StatusOK)
	}
}
