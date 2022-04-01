package authors

import (
	"encoding/json"
	"net/http"
	"sv-article/apps/models"
	"sv-article/databases"
	"sv-article/helpers"
	"sv-article/schemas"

	"github.com/gorilla/mux"
)

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := mux.Vars(r)["id"]

	if models.CheckAuthor(ID) == false {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Author not found")
		return
	}
	var author schemas.Author
	databases.DB.Delete(&author, ID)
	json.NewEncoder(w).Encode("Author has been deleted")
}