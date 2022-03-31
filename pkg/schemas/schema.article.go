package schemas

import "time"

type Article struct {
	ArticleID uint `json:"article_id"`
	Title string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Category string `json:"category" binding:"required"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Status string `json:"status" binding:"required"`
}