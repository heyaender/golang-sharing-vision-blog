package authors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sv-article/apps/models"
	"sv-article/helpers"
	"sv-article/schemas"
)

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	var author schemas.Author
	err := decode.Decode(&author)
	if err != nil {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error Decoding JSON: %s", err.Error()))
		return
	}
	if author.Name == "" {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Name is required")
		return
	}
	if author.Username == "" {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Username is required")
		return
	}
	if author.Password == "" {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Password is required")
		return
	}

	result := models.CreateAuthor(&author)
	if result.Error != nil {
		helpers.JSONErrorResponse(w, http.StatusInternalServerError, "Internal Server Error : can't create author")
		return
	}

	helpers.JSONSuccessResponse(w, author, "Author Created")

}