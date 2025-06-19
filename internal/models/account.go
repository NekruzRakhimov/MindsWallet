package models

import "time"

type Account struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id" db:"user_id"`
	Balance   float64    `json:"balance"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"-" db:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
