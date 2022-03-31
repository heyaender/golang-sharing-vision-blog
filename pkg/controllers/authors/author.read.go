package authors

import (
	"fmt"
	"net/http"
	"sv-article/pkg/helpers"
	"sv-article/pkg/models"

	"github.com/gorilla/mux"
)

func GetAuthorByID (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	fmt.Println(vars)

	author, result := models.GetAuthorByID(ID)
	if result.RowsAffected == 1 {
		helpers.JSONSuccessResponse(w, author, "Author found")
	} else {
		helpers.JSONErrorResponse(w, http.StatusNotFound, "Author not found")
	}
}