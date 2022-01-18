package task

type Storage interface {
	GetOne(uuid string) (*Task, error)
	GetAll(limit, offset int) ([]*Task, error)
	Create(task *Task) error
	Delete(uuid string) error
}
