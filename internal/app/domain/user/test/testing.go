package user

import (
	"demoApi/internal/app/domain/user"
	"testing"
)

func TestValidUser(t *testing.T) *user.User {
	return &user.User{
		Name:     "GOGO",
		Email:    "gogo@example.org",
		Password: "pass",
	}
}

func TestInvalidUser(t *testing.T) *user.User {
	return &user.User{
		Name:     "GOGO",
		Email:    "gogo@example.org",
		Password: "pass",
	}
}

func TestEmptyUser(t *testing.T) *user.User {
	return &user.User{}
}
