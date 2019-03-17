package models

import (
	"github.com/gofrs/uuid"
)

// Subscriptions to repos from users
type Subscriptions struct {
	ID   int       `json:"-" db:"id"`
	User uuid.UUID `json:"user" db:"user"`
	Repo uuid.UUID `json:"repo" db:"repo"`
}
