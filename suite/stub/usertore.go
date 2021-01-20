package stub

import "github.com/ritabc/twg/suite"

// Create variable of type interface, be implemented by type stub.UserStore : compile test
var _ suite.UserStore = &UserStore{}

type UserStore struct{}

func (us *UserStore) Create(user *suite.User) error {
	user.ID = 123
	return nil
}

func (us *UserStore) ByID(id int) (*suite.User, error) {
	return nil, suite.ErrNotFound
}

func (us *UserStore) ByEmail(email string) (*suite.User, error) {
	return nil, nil
}

func (us *UserStore) Delete(*suite.User) error {
	return nil
}
