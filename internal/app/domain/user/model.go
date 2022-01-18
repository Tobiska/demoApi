package user

import "demoApi/internal/app/domain/group"

type User struct {
	Id                int    `json:"id,omitempty"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
	Group             *group.Group
}
