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
	users, err := getUsersWhere("")
	if err != nil {
		return users, fmt.Errorf("unable to get users: %v", err)
	}

	return users, nil
}

// Get returns a single user according to its ID
func (us *Users) Get(id int64) (*User, error) {
	exp := fmt.Sprintf("user_id=%v", id)

	return getUserWhere(exp)
}

// Search returns a single user according to important unique attr
func (us *Users) Search(f string) (*User, error) {
	exp := fmt.Sprintf("username='%v' OR email='%v'", f, f)

	return getUserWhere(exp)
}

// Where return many users according to a map of attrs
func (us *Users) Where(exp string) ([]*User, error) {
	users, err := getUsersWhere(exp)
	if err != nil {
		return nil, fmt.Errorf("unable to get users: %v", err)
	}

	return users, nil
}

// Delete deletes an user based on its ID
func (us *Users) Delete(ID int64) error {
	if err := deleteUser(ID); err != nil {
		return fmt.Errorf("unable to delete user: %v", err)
	}

	return nil
}
