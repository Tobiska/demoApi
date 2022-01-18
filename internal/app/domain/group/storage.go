package group

type Storage interface {
	GetOne(uuid string) (*Group, error)
	GetAll(limit, offset int) ([]*Group, error)
	Create(group *Group) error
	Delete(uuid string) error
}
