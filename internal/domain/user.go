package domain

import "time"

type User struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (*User, error)
	GetByID(id int64) (*User, error)
	UpdateRole(id int64, role string) error
	Delete(id int64) error
	Update(user *User) error
	GetAll() ([]*User, error)
}

type UserService interface {
	GetAllUsers(page, limit int) ([]*User, int, error) // int = total count
	GetUserByID(id int64) (*User, error)
	UpdateUser(user *User) error
	UpdateRole(id int64, role string) error
	DeleteUser(id int64) error
}
