package model

import "time"

type Activity struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
}
