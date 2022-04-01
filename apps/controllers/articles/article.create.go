package articles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sv-article/apps/models"
	"sv-article/helpers"
	"sv-article/schemas"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	decode.DisallowUnknownFields()

	var article schemas.Article
	err := decode.Decode(&article)
	if err != nil {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error Decoding JSON: %s", err.Error()))
		return
	}
	if article.Title == "" {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Title is required")
		return
	}
	if article.Content == "" {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Content is required")
		return
	}
	if article.Category == "" {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Category is required")
		return
	}
	if article.Status == "" {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Status is required")
		return
	}

	result := models.CreateArticle(&article)
	if result.Error != nil {
		helpers.JSONErrorResponse(w, http.StatusInternalServerError, "Internal Server Error : can't create article")
		return
	}

	helpers.JSONSuccessResponse(w, article, "Article Created at "+article.CreatedDate.String())
}