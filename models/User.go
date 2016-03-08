package models

import "github.com/satori/go.uuid"

// User object
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

// NewUser factor method for creating a new User object
func NewUser(name string, username string) User {
	return User{
		ID:       uuid.NewV4().String(),
		Name:     name,
		Username: username,
	}
}

// Equal returns true if the User objects are the same
func (u *User) Equal(other *User) bool {
	if other == nil {
		return false
	}

	return u.ID == other.ID &&
		u.Name == other.Name &&
		u.Username == other.Username
}
