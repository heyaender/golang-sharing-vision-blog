package models

import (
	"sv-article/databases"
	"sv-article/schemas"

	"gorm.io/gorm"
)

func CreateArticle(article *schemas.Article) *gorm.DB {
	result := databases.DB.Create(article)
	return result
}

func GetAllArticle(limit int, offset int) []schemas.Article {
	articles := []schemas.Article{}
	databases.DB.
		Select("id, title, content, category, created_date, updated_date, status, author_name").
		Limit(limit).
		Offset(offset).
		Find(&articles)

		return articles
}

func GetArticleByID(ID string) (schemas.Article, *gorm.DB) {
	var article schemas.Article
	result := databases.DB.
		Select("id, title, content, category, created_date, status, author_id").
		Where("id = ?", ID).
		Find(&article)

	return article, result
}

func GetArticleByStatus(Status string) (schemas.Article, *gorm.DB) {
	var article schemas.Article
	result := databases.DB.
		Select("id, title, content, category, created_date, updated_date, status, author_id").
		Joins("Author").
		Where("status = ?", Status).
		Find(&article)

	return article, result
}

func CheckArticle(ArticleID string) bool {
	var article schemas.Article
	databases.DB.First(&article, ArticleID)
	if article.ID == 0 {
		return false
	}
	return true
}