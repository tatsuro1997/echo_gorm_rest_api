package model

import "time"

type User struct {
	ID        int64     `json:"id" param:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"<-:create"`
	UpdatedAt time.Time `json:"updated_at"`
}
