package domain

import "time"

type EventStaff struct {
	ID        int64      `json:"id" db:"id"`
	EventID   int64      `json:"event_id" db:"event_id"`
	UserID    int64      `json:"user_id" db:"user_id"`
	Role      string     `json:"role" db:"role"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type EventStaffRepository interface {
	Create(eventStaff *EventStaff) error
	GetByID(id int64) (*EventStaff, error)
	GetByEventID(eventID int64) ([]*EventStaff, error)
	GetByUserID(userID int64) ([]*EventStaff, error)
	Update(eventStaff *EventStaff) error
	Delete(id int64) error
}

type EventStaffService interface {
	AssignStaff(eventID, userID int64, role string) error
	RevokeStaff(eventID, userID int64) error
	UpdateStaffRole(eventID, userID int64, role string) error
	GetStaffByEventID(eventID int64) ([]*EventStaff, error)
	GetStaffByUserID(userID int64) ([]*EventStaff, error)
	GetStaffRole(eventID, userID int64) (string, bool, error)
}
