package model

type User struct {
	Model
	Name  string `json:"name"`
	Posts []Post `json:"posts" gorm:"foreignKey:UserID" param:"user_id"`
}
