package user

type Storage interface {
	GetOne(uuid string) (*User, error)
	GetAll(limit, offset int) ([]*User, error)
	Create(user *User) error
	Delete(uuid string) error
}
