package models

import (
	"github.com/gofrs/uuid"
)

// Subscriptions to repos from users
type Subscriptions struct {
	UserID uuid.UUID `gorm:"user_id" json:"user"`
	RepoID uuid.UUID `gorm:"repo_id" json:"repo" `
}
