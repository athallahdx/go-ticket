package domain

import "time"

type Order struct {
	ID               int64      `json:"id" db:"id"`
	UserID           int64      `json:"user_id" db:"user_id"`
	TotalAmount      float64    `json:"total_amount" db:"total_amount"`
	Status           string     `json:"status" db:"status"`
	PaymentMethod    string     `json:"payment_method" db:"payment_method"`
	PaymentReference string     `json:"payment_reference" db:"payment_reference"`
	ExpiredAt        time.Time  `json:"expired_at" db:"expired_at"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	DeletedAt        *time.Time `json:"deleted_at" db:"deleted_at"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id int64) (*Order, error)
	GetByUserID(userID int64) ([]*Order, error)
	Update(order *Order) error
	Delete(id int64) error
}

type OrderService interface {
	CreateOrder(order *Order) error
	GetOrderByID(id int64) (*Order, error)
	GetOrdersByUserID(userID int64) ([]*Order, error)
	UpdateOrder(order *Order) error
	DeleteOrder(id int64) error
}
