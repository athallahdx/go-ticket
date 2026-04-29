// domain/checkin.go
package domain

import "time"

type Checkin struct {
	ID          int64     `json:"id" db:"id"`
	TicketID    int64     `json:"ticket_id" db:"ticket_id"`
	CheckedInAt time.Time `json:"checked_in_at" db:"checked_in_at"`
	CheckedInBy int64     `json:"checked_in_by" db:"checked_in_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type CheckinStats struct {
	TotalTickets   int
	TotalCheckedIn int
	TotalRemaining int
	ByTicketType   map[string]int
}

type CheckinRepository interface {
	Create(checkin *Checkin) error
	GetByID(id int64) (*Checkin, error)
	GetByTicketID(ticketID int64) (*Checkin, error)
	GetByEventID(eventID int64, page, limit int) ([]*Checkin, int, error)
	GetByStaffID(staffID int64) ([]*Checkin, error)
}

type CheckinService interface {
	CheckinTicket(ticketCode string, staffID int64) (*Checkin, error)
	ValidateTicket(ticketCode string) (*Ticket, error)
	GetCheckinHistory(eventID int64, page, limit int) ([]*Checkin, int, error)
	GetCheckinStats(eventID int64) (*CheckinStats, error)
}
