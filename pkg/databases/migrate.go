package databases

import (
	"fmt"
	"sv-article/pkg/schemas"
)

func DBMigrate() {
	DB.AutoMigrate(&schemas.Author{}, &schemas.Article{})
	fmt.Println("DB Migration Success")
}