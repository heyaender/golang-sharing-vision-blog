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
	ArticleID := vars["article_id"]

	fmt.Println(vars)

	article, result := models.GetArticleByID(ArticleID)
	if result.RowsAffected == 1 {
		helpers.JSONSuccessResponse(w, article, "Article found")
	} else {
		helpers.JSONErrorResponse(w, http.StatusNotFound, "Article not found")
	}
}