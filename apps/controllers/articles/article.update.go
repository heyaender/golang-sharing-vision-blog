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

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]

	if models.CheckArticle(ID) == false {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Article not found")
		return
	}

	var article schemas.Article

	databases.DB.First(&article, ID)
	json.NewDecoder(r.Body).Decode(&article)
	databases.DB.Save(&article)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}