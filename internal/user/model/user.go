package model

import "time"

type User struct {
	ID           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdateAt     time.Time `json:"update_at"`
}
