package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofrs/uuid"
)

// User they do things
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Email     string    `json:"email" db:"email"`
	Githubid  string    `json:"github_id" db:"github_id"`
}

// Users are groups of people that do things
type Users []User

// BeforeCreate is called from GORM
func (u *User) BeforeCreate() (err error) {
	if u.ID, err = uuid.NewV4(); err != nil {
		return err
	}

	return validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Githubid, validation.Required),
	)
}
