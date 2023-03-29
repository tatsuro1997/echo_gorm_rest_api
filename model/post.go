package model

type Post struct {
	Model
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id" gorm:"not null"`
}
