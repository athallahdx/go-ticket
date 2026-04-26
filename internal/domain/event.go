package domain

import "time"

type Event struct {
	ID          int64        `json:"id" db:"id"`
	OrganizerID int64        `json:"organizer_id" db:"organizer_id"`
	Thumbnail   string       `json:"thumbnail" db:"thumbnail"`
	Images      []EventImage `json:"images"`
	Name        string       `json:"name" db:"name"`
	Description string       `json:"description" db:"description"`
	Location    string       `json:"location" db:"location"`
	Date        time.Time    `json:"date" db:"date"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt   time.Time    `json:"deleted_at" db:"deleted_at"`
}

type EventRepository interface {
	Create(event *Event) error
	GetByID(id int64) (*Event, error)
	GetAll() ([]*Event, error)
	Update(event *Event) error
	Delete(id int64) error
}

type EventFilter struct {
	OrganizerID int64
	City        string
	DateFrom    time.Time
	DateTo      time.Time
	Search      string
}

type EventService interface {
	CreateEvent(event *Event) error
	UpdateEvent(event *Event) error
	DeleteEvent(id int64) error
	GetEventByID(id int64) (*Event, error)
	GetAllEvents(filter EventFilter, page, limit int) ([]*Event, int, error)
	GetEventsByOrganizerID(organizerID int64, page, limit int) ([]*Event, int, error)

	AddImage(eventID int64, imageURL string) (*EventImage, error)
	RemoveImage(imageID int64) error
	GetImagesByEventID(eventID int64) ([]*EventImage, error)
}
