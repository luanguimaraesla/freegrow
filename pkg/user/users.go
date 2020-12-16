package user

import "fmt"

// Users is the entity used for controlling
// interactions with many users in the database
type Users struct{}

// NewUsers returns an Users controller
func NewUsers() *Users {
	return &Users{}
}

// All returns all the users registered in the database
func (us *Users) All() ([]*User, error) {
	users, err := getAllUsers()
	if err != nil {
		return users, fmt.Errorf("unable to get users: %v", err)
	}

	return users, nil
}

// Get returns the user information according to its ID
func (us *Users) Get(ID int64) (*User, error) {
	user, err := getUserByID(ID)
	if err != nil {
		return nil, fmt.Errorf("unable to load user: %v", err)
	}

	return user, nil
}

// Delete deletes an user based on its ID
func (us *Users) Delete(ID int64) error {
	if err := deleteUser(ID); err != nil {
		return fmt.Errorf("unable to delete user: %v", err)
	}

	return nil
}
