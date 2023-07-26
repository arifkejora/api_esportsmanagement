package model

import "time"

type MAttendance struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	EventID   int       `json:"event_id"`
	TimeStamp time.Time `json:"timestamp"`
}

type attendance struct {
	UserID  int `json:"user_id"`
	EventID int `json:"event_id"`
}
