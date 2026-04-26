// -- ======================
// -- ORGANIZERS
// -- ======================
// CREATE TABLE IF NOT EXISTS organizers (
//     id BIGINT AUTO_INCREMENT PRIMARY KEY,
//     user_id BIGINT NOT NULL,
//     company_name VARCHAR(255) NOT NULL,
//     phone VARCHAR(20),
//     email VARCHAR(100),
//     logo VARCHAR(255),
//     description TEXT,
//     city VARCHAR(100),
//     province VARCHAR(100),
//     is_verified BOOLEAN DEFAULT FALSE,
//     verified_at TIMESTAMP NULL DEFAULT NULL,
//     verified_by BIGINT NULL,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//     deleted_at TIMESTAMP NULL DEFAULT NULL,
//     FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
//     FOREIGN KEY (verified_by) REFERENCES users(id) ON DELETE SET NULL,
//     INDEX idx_organizers_deleted_at (deleted_at),
//     INDEX idx_organizers_is_verified (is_verified)
// );

package domain

import "time"

type Organizer struct {
	ID          int64     `json:"id" db:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	CompanyName string    `json:"company_name" db:"company_name"`
	Phone       string    `json:"phone" db:"phone"`
	Email       string    `json:"email" db:"email"`
	Logo        string    `json:"logo" db:"logo"`
	Description string    `json:"description" db:"description"`
	City        string    `json:"city" db:"city"`
	Province    string    `json:"province" db:"province"`
	IsVerified  bool      `json:"is_verified" db:"is_verified"`
	VerifiedAt  time.Time `json:"verified_at" db:"verified_at"`
	VerifiedBy  int64     `json:"verified_by" db:"verified_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at" db:"deleted_at"`
}

type OrganizerRepository interface {
	Create(organizer *Organizer) error
	GetByID(id int64) (*Organizer, error)
	GetByUserID(userID int64) (*Organizer, error)
}

type OrganizerService interface {
	RegisterAsOrganizer(user *User, companyName string) (*Organizer, error)
	GetOrganizerByID(id int64) (*Organizer, error)
	GetOrganizerByUserID(userID int64) (*Organizer, error)
	GetAllOrganizers(page, limit int) ([]*Organizer, int, error)
	UpdateOrganizer(organizer *Organizer) error
	DeleteOrganizer(id int64) error

	VerifyOrganizer(organizerID int64, adminID int64) error
	RevokeVerification(organizerID int64, adminID int64) error
}
