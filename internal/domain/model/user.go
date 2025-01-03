package model

import "time"

type User struct {
	Id        int       `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Name      string    `json:"name" gorm:"not sssnull"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(name string) *User {
	return &User{
		Name: name,
	}
}
