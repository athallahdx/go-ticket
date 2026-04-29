package domain

import "time"

type Ticket struct {
	ID           int64      `json:"id" db:"id"`
	Code         string     `json:"code" db:"code"`
	UserID       int64      `json:"user_id" db:"user_id"`
	TicketTypeID int64      `json:"ticket_type_id" db:"ticket_type_id"`
	QRCode       string     `json:"qr_code" db:"qr_code"`
	Status       string     `json:"status" db:"status"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}

type TicketRepository interface {
	Create(ticket *Ticket) error
	GetByID(id int64) (*Ticket, error)
	GetByCode(code string) (*Ticket, error)
	GetByUserID(userID int64) ([]*Ticket, error)
	GetByTicketTypeID(ticketTypeID int64) ([]*Ticket, error)
	Update(ticket *Ticket) error
	Delete(id int64) error
}

type TicketService interface {
	CreateTicket(ticket *Ticket) error
	GetTicketByID(id int64) (*Ticket, error)
	GetTicketByCode(code string) (*Ticket, error)
	GetTicketsByUserID(userID int64) ([]*Ticket, error)
	GetTicketsByTicketTypeID(ticketTypeID int64) ([]*Ticket, error)
	UpdateTicket(ticket *Ticket) error
	DeleteTicket(id int64) error
}
