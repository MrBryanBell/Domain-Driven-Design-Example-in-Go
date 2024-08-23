package UserUpdater

import User "mymodule/User/Domain"

type UserUpdater struct {
	repository User.UserRepository
}

func Setup(repository User.UserRepository) UserUpdater {
	return UserUpdater{repository: repository}
}

func (self *UserUpdater) Run(id, userName, newEmail string) (*User.User, error) {

	user, err := self.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	updatedUser, err := user.UpdateEmail(newEmail)
	if err != nil {
		return nil, err
	}

	err = self.repository.Save(&updatedUser)
	if err != nil {
		return nil, err
	}

	return &updatedUser, nil
}
