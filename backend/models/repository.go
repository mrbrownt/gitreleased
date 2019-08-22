package models

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Repository currnt matches github repo
type Repository struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"-" db:"created_at"`
	UpdatedAt   time.Time `json:"-" db:"updated_at"`
	Owner       string    `json:"owner,omitempty" db:"owner"`
	Name        string    `json:"name,omitempty" db:"name"`
	Description string    `json:"description,omitempty" db:"description"`
	URL         string    `json:"url,omitempty" db:"url"`
}

// Exists checks if the record exists
func (r *Repository) Exists(db *gorm.DB) bool {
	db.Where(r).First(&r)
	return r.ID != uuid.Nil
}
