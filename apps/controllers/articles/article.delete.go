package articles

import (
	"encoding/json"
	"net/http"
	"sv-article/apps/models"
	"sv-article/databases"
	"sv-article/helpers"
	"sv-article/schemas"

	"github.com/gorilla/mux"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := mux.Vars(r)["id"]

	if models.CheckArticle(ID) == false {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Article not found")
		return
	}

	var article schemas.Article
	databases.DB.Delete(&article, ID)
	json.NewEncoder(w).Encode("Article has been deleted")
}