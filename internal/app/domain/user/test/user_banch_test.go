package user

import (
	"demoApi/internal/app/domain/user"
	"testing"
)

func BenchmarkEncryptPassword(b *testing.B) {
	u := &user.User{
		Name:     "GOGO",
		Email:    "gogo@mail.ru",
		Password: "1q2w3e",
	}
	for i := 0; i < b.N; i++ {
		err := u.EncryptPassword()
		if err != nil {
			return
		}
	}
}
