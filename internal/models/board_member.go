package models

import "time"

type BoardMember struct {
	ID        int       `json:"id" db:"id"`
	BoardID   int       `json:"board_id" db:"board_id"`
	UserID    int       `json:"user_id" db:"user_id"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type BoardMemberRequest struct {
	UserID int    `json:"user_id" validate:"required"`
	Role   string `json:"role" validate:"required,oneof=owner admin member"`
}