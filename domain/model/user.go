package model

import (
	"time"
)

// User is a user of the system.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Furigana  string    `json:"furigana"`
	Photo     string    `json:"photo"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
}
