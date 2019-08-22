package models

import (
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// Subscriptions to repos from users
type Subscriptions struct {
	UserID uuid.UUID `gorm:"user_id" json:"user"`
	RepoID uuid.UUID `gorm:"repo_id" json:"repo" `
}

func (s *Subscriptions) Create(db *gorm.DB) (err error) {
	if !s.Valid() {
		return errors.New("subscription is not valid")
	}
	return db.Where(s).FirstOrCreate(&s).Error
}

func (s *Subscriptions) Valid() bool {
	if s.RepoID == uuid.Nil {
		return false
	}

	if s.UserID == uuid.Nil {
		return false
	}
	return true
}
