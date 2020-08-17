package models

import (
	"errors"
	"fmt"
)

// User data
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers returns the collection of existing users
func GetUsers() []*User {
	return users
}

// AddUser adds the provided User to the users collection
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("User must not include ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)

	return u, nil
}

// GetUserByID returns a user matching the provided ID (if exists)
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser does exactly that
func UpdateUser(u User) (User, error) {
	// ToDo maybe GetUserByID can be used here
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID removes a user...by ID
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}
