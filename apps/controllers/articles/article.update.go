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
	ArticleID := mux.Vars(r)["article_id"]

	if models.CheckArticle(ArticleID) == false {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Article not found")
		return
	}

	var article schemas.Article

	databases.DB.First(&article, ArticleID)
	json.NewDecoder(r.Body).Decode(&article)
	databases.DB.Save(&article)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}