package user

import (
	"demoApi/internal/app/domain/group"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int          `json:"id,omitempty"`
	Name              string       `json:"name"`
	Email             string       `json:"email"`
	Password          string       `json:"password,omitempty"`
	EncryptedPassword string       `json:"-"`
	Group             *group.Group `json:"group,omitempty"`
}

func (u *User) EncryptPassword() error {
	enc, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}
	u.EncryptedPassword = string(enc)
	u.sanitize()
	return err
}

func (u *User) sanitize() {
	u.Password = ""
}
