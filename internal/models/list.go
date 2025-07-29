package models

import "time"

type List struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	BoardID   int       `json:"board_id" db:"board_id"`
	Position  int       `json:"position" db:"position"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ListRequest struct {
	Title string `json:"title" validate:"required,min=1,max=100"`
}