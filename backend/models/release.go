package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Release of a repository
type Release struct {
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	RepositoryID uuid.UUID `json:"repository_id" db:"repository_id"`
	Version      string    `json:"version" db:"version"`
	Major        int       `json:"major" db:"major"`
	Minor        int       `json:"minor" db:"minor"`
	Patch        int       `json:"patch" db:"patch"`
	Development  string    `json:"development" db:"development"`
}
