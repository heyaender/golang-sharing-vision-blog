package routes

import (
	"sv-article/pkg/controllers"
	"sv-article/pkg/controllers/authors"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	Route := mux.NewRouter().StrictSlash(true)
	Route.HandleFunc("/", controllers.Index)
	Route.HandleFunc("/author", authors.CreateAuthor).Methods("OPTIONS", "POST")
	Route.HandleFunc("/authors", authors.GetAllAuthor).Methods("OPTIONS", "GET")
	Route.HandleFunc("/author/{id}", authors.GetAuthorByID).Methods("OPTIONS", "GET")
	Route.HandleFunc("/author/{id}", authors.UpdateAuthor).Methods("OPTIONS", "PUT")
	Route.HandleFunc("/author/{id}", authors.DeleteAuthor).Methods("OPTIONS", "DELETE")

	return Route
}