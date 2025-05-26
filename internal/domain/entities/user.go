package entities

import (
	"errors"
	"time"

	"github.com/felipegonrod/sophos/internal/domain/valueobjects"
)

type UserRole string

const (
	RoleReader UserRole = "reader"
	RoleAuthor UserRole = "author"
	RoleAdmin  UserRole = "admin"
)

type User struct {
	id        string
	email     *valueobjects.Email // Now using the value object!
	username  string
	role      UserRole
	createdAt time.Time
	updatedAt time.Time
}

// Constructor now validates email through value object
func NewUser(id, emailStr, username string) (*User, error) {
	// Email validation is now handled by the value object
	email, err := valueobjects.NewEmail(emailStr)
	if err != nil {
		return nil, err // Email validation error bubbles up
	}

	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	return &User{
		id:        id,
		email:     email,
		username:  username,
		role:      RoleReader,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

// Getters
func (u *User) ID() string                 { return u.id }
func (u *User) Email() *valueobjects.Email { return u.email }
func (u *User) EmailString() string        { return u.email.String() } // Convenience method
func (u *User) Username() string           { return u.username }
func (u *User) Role() UserRole             { return u.role }
func (u *User) CreatedAt() time.Time       { return u.createdAt }

// Business methods can now use email domain logic
func (u *User) IsFromDomain(domain string) bool {
	return u.email.Domain() == domain
}

func (u *User) PromoteToAuthor() error {
	if u.role == RoleAdmin {
		return errors.New("admin cannot be demoted to author")
	}
	u.role = RoleAuthor
	u.updatedAt = time.Now()
	return nil
}

func (u *User) CanCreatePost() bool {
	return u.role == RoleAuthor || u.role == RoleAdmin
}
