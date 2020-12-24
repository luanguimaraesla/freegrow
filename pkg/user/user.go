package user

import (
	"fmt"

	"github.com/luanguimaraesla/freegrow/pkg/gadget"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// New returns an empty user
func New() *User {
	return &User{}
}

// Create inserts the user in the database
func (u *User) Create() error {
	if err := insertUser(u); err != nil {
		return fmt.Errorf("unable to create user: %v", err)
	}

	return nil
}

// Update updates the user record in the database
func (u *User) Update() error {
	if err := updateUser(u); err != nil {
		return fmt.Errorf("unable to update user", err)
	}

	return nil
}

// Delete deletes the user record from the database
func (u *User) Delete() error {
	if err := deleteUser(u.ID); err != nil {
		return fmt.Errorf("unable to delete user", err)
	}

	return nil
}

// Gadgets returns a list of user's gadgets
func (u *User) Gadgets() *gadget.Gadgets {
	return gadget.NewGadgets(u.ID)
}

// checkPassword receives a plain password and compares
// with the encrypted one stored in the database
func (u *User) checkPassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return false
	}

	return true
}
