package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Repository currnt matches github repo
type Repository struct {
	ID          uuid.UUID `json:"id,omitempty" db:"id"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Owner       string    `json:"owner,omitempty"`
	Name        string    `json:"name,omitempty"`
	GithubID    string    `json:"github_id,omitempty"`
	Description string    `json:"description,omitempty"`
	Releases    []Release `json:"releases,omitempty"`
}
