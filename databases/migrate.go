package databases

import (
	"fmt"
	"sv-article/schemas"
)

func DBMigrate() {
	DB.AutoMigrate(&schemas.Author{}, &schemas.Article{})
	fmt.Println("DB Migration Success")
}
