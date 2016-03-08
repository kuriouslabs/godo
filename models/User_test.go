package models

import "testing"

func TestNewUser_generatesID(t *testing.T) {
	name := "first last"
	username := "username"
	user := NewUser(name, username)

	if user.ID == "" {
		t.Error("Failed to generate UUID for user")
	}
}
