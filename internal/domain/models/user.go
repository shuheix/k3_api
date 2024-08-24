package models

import "time"

type User struct {
	Id   int    `json:"id" gorm:"primaryKey" gorm:"autoIncrement"`
	Name string `json:"name" gorm:"not_null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}

type CreateUserParams struct {
	Name string `json:"name"`
}
