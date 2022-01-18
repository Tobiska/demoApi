package session

type Storage interface {
	GetOne(uuid string) (*Session, error)
	GetAll(limit, offset int) ([]*Session, error)
	Create(session *Session) error
	Delete(uuid string) error
}
