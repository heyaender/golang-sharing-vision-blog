package schemas

import "time"

type Article struct {
	ID int `json:"id"`
	AuthorID uint `json:"author_id"`
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required" gorm:"type:text" `
	Category string `json:"category" binding:"required"`
	CreatedDate time.Time `json:"-" gorm:"type:timestamp"`
	UpdatedDate time.Time `json:"-" gorm:"type:timestamp"`
	Status string `json:"status" binding:"required"`
}