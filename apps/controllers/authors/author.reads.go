package authors

import (
	"fmt"
	"net/http"
	"strconv"
	"sv-article/apps/models"
	"sv-article/helpers"
)

func GetAllAuthor(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 0

	limitQuery := r.URL.Query().Get("limit")
	offsetQuery := r.URL.Query().Get("offset")

	var err error
	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			helpers.JSONErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error parsing limit query: %s", err.Error()))
			return
		}
	}
	if offsetQuery != "" {
		offset, err = strconv.Atoi(offsetQuery)
		if err != nil {
			helpers.JSONErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("Error parsing offset query: %s", err.Error()))
			return
		}
	}

	authors := models.GetAllAuthor(limit, offset)
	helpers.JSONSuccessResponse(w, authors, "Successfully get all authors")
}