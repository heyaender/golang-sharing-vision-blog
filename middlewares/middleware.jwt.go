package middlewares

import (
	"net/http"
	"strings"
	"sv-article/helpers"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) < 2 {
			helpers.JSONErrorResponse(w, http.StatusForbidden, "Invalid Token")
			return
		}

		token := bearerToken[1]

		_, err := helpers.VerifyJwtToken(token)
		if err != nil {
			helpers.JSONErrorResponse(w, http.StatusForbidden, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}