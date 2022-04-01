package models

import (
	"sv-article/databases"
	"sv-article/schemas"

	"gorm.io/gorm"
)

func CreateAuthor(author *schemas.Author) *gorm.DB {
	result := databases.DB.Create(author)
	return result
}

func GetAllAuthor(limit int, offset int) []schemas.Author {
	authors := []schemas.Author{}
	databases.DB.
		Select("id, name, username, password").
		Limit(limit).
		Offset(offset).
		Find(&authors)

		return authors
}

func GetAuthorByID(ID string) (schemas.Author, *gorm.DB) {
	var authors schemas.Author
	result := databases.DB.
		Select("authors.id, authors.name, authors.username, authors.password").
		// Joins("Article").
		Where("authors.id = ?", ID).
		Find(&authors)
	
	return authors, result
}

func CheckAuthor(ID string) bool {
	var author schemas.Author
	databases.DB.First(&author, ID)
	if author.ID == 0 {
		return false
	}
	return true
}
