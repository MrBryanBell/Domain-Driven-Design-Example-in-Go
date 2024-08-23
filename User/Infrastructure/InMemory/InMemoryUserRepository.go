package InMemoryUserRepository

import User "mymodule/User/Domain"
import "errors"

type InMemoryUserRepository struct {
	users []User.User
}


func New()*InMemoryUserRepository {
	return &InMemoryUserRepository{users: []User.User{}}
}

func (self *InMemoryUserRepository) FindByID(id string) (*User.User, error) {
	for i, user := range self.users {
		if user.ID == id {
			return &self.users[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func (self *InMemoryUserRepository) Save(user *User.User) error {
	for i := range self.users {
		if self.users[i].ID == user.ID {
			self.users[i] = *user
			return nil
		}
	}
	self.users = append(self.users, *user)
	return nil
}
