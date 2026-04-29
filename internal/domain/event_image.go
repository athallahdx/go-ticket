package domain

import "time"

type EventImage struct {
	ID        int64     `json:"id" db:"id"`
	EventID   int64     `json:"event_id" db:"event_id"`
	ImageURL  string    `json:"image_url" db:"image_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
