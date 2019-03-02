package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofrs/uuid"
)

// User they do things
type User struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Email          string    `json:"email" db:"email"`
	GithubID       string    `json:"github_id" db:"github_id"`
	GithubUserName string    `json:"github_user_name" db:"github_user_name"`
	FirstName      string    `json:"first_name" db:"first_name"`
	LastName       string    `json:"last_name" db:"last_name"`
}

// Users are groups of people that do things
type Users []User

// BeforeCreate is called from GORM
func (u *User) BeforeCreate() (err error) {
	if u.ID, err = uuid.NewV4(); err != nil {
		return err
	}

	return validation.ValidateStruct(u,
		validation.Field(&u.Email, is.Email),
		validation.Field(&u.GithubID, validation.Required),
	)
}
