package routes

import (
	"sv-article/apps/controllers"
	"sv-article/apps/controllers/articles"
	"sv-article/apps/controllers/authors"

	"github.com/gorilla/mux"
)

func Route() *mux.Router {
	Route := mux.NewRouter().StrictSlash(true)
	Route.HandleFunc("/", controllers.Index)

	// Author Route
	Route.HandleFunc("/author", authors.CreateAuthor).Methods("OPTIONS", "POST")
	Route.HandleFunc("/authors", authors.GetAllAuthor).Methods("OPTIONS", "GET")
	Route.HandleFunc("/author/{id}", authors.GetAuthorByID).Methods("OPTIONS", "GET")
	Route.HandleFunc("/author/{id}", authors.UpdateAuthor).Methods("OPTIONS", "PUT")
	Route.HandleFunc("/author/{id}", authors.DeleteAuthor).Methods("OPTIONS", "DELETE")

	// Article Route
	Route.HandleFunc("/article", articles.CreateArticle).Methods("OPTIONS", "POST")
	Route.HandleFunc("/articles", articles.GetAllArticle).Methods("OPTIONS", "GET")
	Route.HandleFunc("/article/{id}", articles.GetArticleByID).Methods("OPTIONS", "GET")
	Route.HandleFunc("/article/{id}", articles.UpdateArticle).Methods("OPTIONS", "PUT")
	Route.HandleFunc("/article/{id}", articles.DeleteArticle).Methods("OPTIONS", "DELETE")
	Route.HandleFunc("/articles_status/{status}", articles.GetArticleByStatus).Methods("OPTIONS", "GET")

	return Route
}