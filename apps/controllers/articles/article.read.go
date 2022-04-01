package articles

import (
	"fmt"
	"net/http"
	"sv-article/apps/models"
	"sv-article/helpers"

	"github.com/gorilla/mux"
)

func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	fmt.Println(vars)

	article, result := models.GetArticleByID(ID)
	if result.RowsAffected == 1 {
		helpers.JSONSuccessResponse(w, article, "Article found")
	} else {
		helpers.JSONErrorResponse(w, http.StatusNotFound, "Article not found")
	}
}

func GetArticleByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]

	articles, result := models.GetArticleByStatus(status)
	if result.RowsAffected > 0 {
		helpers.JSONSuccessResponse(w, articles, "Article found")
	} else {
		helpers.JSONErrorResponse(w, http.StatusNotFound, "Article not found")
	}
}