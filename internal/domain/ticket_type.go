package domain

import "time"

type TicketType struct {
	ID          int64      `json:"id" db:"id"`
	EventID     int64      `json:"event_id" db:"event_id"`
	Name        string     `json:"name" db:"name"`
	Price       float64    `json:"price" db:"price"`
	Quota       int        `json:"quota" db:"quota"`
	Sold        int        `json:"sold" db:"sold"`
	Description string     `json:"description" db:"description"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

type TicketTypeRepository interface {
	Create(ticketType *TicketType) error
	GetByID(id int64) (*TicketType, error)
	GetByEventID(eventID int64) ([]*TicketType, error)
	Update(ticketType *TicketType) error
	Delete(id int64) error
}

type TicketTypeService interface {
	CreateTicketType(ticketType *TicketType) error
	UpdateTicketType(ticketType *TicketType) error
	DeleteTicketType(id int64) error
	GetTicketTypeByID(id int64) (*TicketType, error)
	GetTicketTypesByEventID(eventID int64) ([]*TicketType, error)
	GetAvailableTicketTypes(eventID int64) ([]*TicketType, error)
}
