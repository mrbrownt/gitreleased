package models

import (
	"time"

	"github.com/gofrs/uuid"
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
