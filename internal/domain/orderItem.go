package domain

import "time"

type OrderItem struct {
	ID           int64     `json:"id" db:"id"`
	OrderID      int64     `json:"order_id" db:"order_id"`
	TicketTypeID int64     `json:"ticket_type_id" db:"ticket_type_id"`
	Quantity     int       `json:"quantity" db:"quantity"`
	Price        float64   `json:"price" db:"price"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	DeletedAt    time.Time `json:"deleted_at" db:"deleted_at"`
}

type OrderItemRepository interface {
	Create(orderItem *OrderItem) error
	GetByID(id int64) (*OrderItem, error)
	GetByOrderID(orderID int64) ([]*OrderItem, error)
	GetByTicketTypeID(ticketTypeID int64) ([]*OrderItem, error)
	Update(orderItem *OrderItem) error
	Delete(id int64) error
}

type OrderItemService interface {
	CreateOrderItem(orderItem *OrderItem) error
	GetOrderItemByID(id int64) (*OrderItem, error)
	GetOrderItemsByOrderID(orderID int64) ([]*OrderItem, error)
	GetOrderItemsByTicketTypeID(ticketTypeID int64) ([]*OrderItem, error)
	UpdateOrderItem(orderItem *OrderItem) error
	DeleteOrderItem(id int64) error
}
