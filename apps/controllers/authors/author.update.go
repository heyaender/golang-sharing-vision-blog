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

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	ID := mux.Vars(r)["id"]
	
	if models.CheckAuthor(ID) == false {
		helpers.JSONErrorResponse(w, http.StatusBadRequest, "Author not found")
		return
	}

	var author schemas.Author

	databases.DB.First(&author, ID)
	json.NewDecoder(r.Body).Decode(&author)
	databases.DB.Save(&author)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}