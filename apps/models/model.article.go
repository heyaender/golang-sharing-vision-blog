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
		Select("aritcle_id, title, content, category, created_date, updated_date, status").
		Limit(limit).
		Offset(offset).
		Find(&articles)

		return articles
}

func GetArticleByID(ArticleID string) (schemas.Article, *gorm.DB) {
	var article schemas.Article
	result := databases.DB.
		Select("aritcle_id, title, content, category, created_date, updated_date, status").
		Joins("Author").
		Where("aritcle_id = ?", ArticleID).
		Find(&article)

	return article, result
}

func CheckArticle(ArticleID string) bool {
	var article schemas.Article
	databases.DB.First(&article, ArticleID)
	if article.ArticleID == 0 {
		return false
	}
	return true
}