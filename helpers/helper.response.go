package helpers

import (
	"encoding/json"
	"net/http"
	"sv-article/schemas"
)

func JSONSuccessResponse(w http.ResponseWriter, Data interface{}, message string) {
	result, _ := json.Marshal(schemas.Response{
		Code : http.StatusOK,
		Data : Data,
		Message : message,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func JSONErrorResponse(w http.ResponseWriter, StatusCode int, message string) {
	result, _ := json.Marshal(schemas.Response{
		Code : StatusCode,
		Message : message,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusCode)
	w.Write(result)
}