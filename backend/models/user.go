package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

// User they do things
type User struct {
	ID             uuid.UUID `json:"id,omitempty" db:"id"`
	CreatedAt      time.Time `json:"-" db:"created_at"`
	UpdatedAt      time.Time `json:"-" db:"updated_at"`
	Email          string    `json:"email,omitempty" db:"email"`
	GithubID       string    `json:"github_id,omitempty" db:"github_id"`
	GithubUserName string    `json:"github_user_name,omitempty" db:"github_user_name"`
	FirstName      string    `json:"first_name,omitempty" db:"first_name"`
	LastName       string    `json:"last_name,omitempty" db:"last_name"`
	AccessToken    string    `json:"access_token,omitempty" ab:"access_token"`
}

// Users are groups of people that do things
type Users []User

// BeforeCreate is called from GORM
func (u *User) BeforeCreate() (err error) {
	return validation.ValidateStruct(u,
		validation.Field(&u.Email, is.Email),
		validation.Field(&u.GithubID, validation.Required),
	)
}

func (u *User) Valid(db *gorm.DB) bool {
	db.Where(u).First(&u)
	return u.ID != uuid.Nil
}
