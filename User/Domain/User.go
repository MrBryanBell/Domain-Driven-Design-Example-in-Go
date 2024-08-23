/*
	✌️USER NAME
	-[x] Name cannot be empty
	-[x] Name should have a minimum 3 characters
	-[x] Name should have a maximum 50 characters

	✌️USER EMAIL
	-[x] Email cannot be empty
	-[x] Email should be gmail
*/

/*
	MAKE IT WORK,
	MAKE IT RIGHT,
	MAKE IT FAST
*/

package User

import (
	"mymodule/User/Domain/UserEmail"
	"mymodule/User/Domain/UserName"
)

type User struct {
	ID string
	name  UserName.UserName
	email UserEmail.UserEmail
}

func Create(id, name, email string) (User, error) {

	_name, error := UserName.New(name)
	if error != nil {
		return User{}, error
	}

	_email, error := UserEmail.New(email)
	if error != nil {
		return User{}, error
	}

	return User{
		ID: id,
		name:  _name,
		email: _email,
	}, nil

}

func (user *User) UpdateEmail(newEmail string) (User, error) {
	_newEmail, error := UserEmail.New(newEmail)

	if error != nil {
		return User{}, error
	}

	return User{
		ID: user.ID,
		name:  user.name,
		email: _newEmail,
	}, nil
}