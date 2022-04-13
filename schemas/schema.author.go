package schemas

type Author struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Article *Article `gorm:"foreignkey:AuthorID" json:"article,omitempty"`
	// Article []Article `gorm:"foreignkey:AuthorID;association_foreignkey:ID" json:"article"`
}