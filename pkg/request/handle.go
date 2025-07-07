package request

import (
	"go/http/pkg/response"
	"net/http"
)

// Попробовать убрать указатель на интерфейс w, т.к он не нужен
func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {

	body, err := Decode[T](r.Body)
	if err != nil {
		response.Json(w, err.Error(), 402)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		response.Json(w, err.Error(), 402)
		return nil, err
	}

	return &body, nil

}
