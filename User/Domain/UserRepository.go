package User

type UserRepository interface {
	Save(user User) error
}
