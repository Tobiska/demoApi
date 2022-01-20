package user_test

import (
	"context"
	user2 "demoApi/internal/app/adapters/db/inmemory/user"
	"demoApi/internal/app/domain/user"
	user3 "demoApi/internal/app/domain/user/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_GetByID(t *testing.T) {
	testCases := []struct {
		name    string
		mp      map[int]*user.User
		uid     int
		isValid bool
	}{
		{
			name: "valid user",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       1,
					Name:     "BLA",
					Password: "pass",
				},
			},
			uid:     1,
			isValid: true,
		},

		{
			name: "notfound user",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       1,
					Name:     "BLA",
					Password: "pass",
				},
			},
			uid:     2,
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := user.NewService(user2.NewStorage(&tc.mp)).GetByID(context.TODO(), tc.uid)
			if tc.isValid {
				assert.NoError(t, err)
				assert.Equal(t, tc.uid, u.Id)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUser_GetByAll(t *testing.T) {
	testCases := []struct {
		name        string
		mp          map[int]*user.User
		offset      int
		limit       int
		validResult []int
	}{
		{
			name: "full users",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       1,
					Name:     "BLA",
					Password: "pass",
				},
			},
			offset:      0,
			limit:       2,
			validResult: []int{0, 1},
		},

		{
			name:        "empty users",
			mp:          map[int]*user.User{},
			offset:      0,
			limit:       0,
			validResult: []int{},
		},

		{
			name: "range less users",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       1,
					Name:     "BLA",
					Password: "pass",
				},
			},
			offset:      0,
			limit:       1,
			validResult: []int{0},
		},

		{
			name: "range grater users",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       1,
					Name:     "BLA",
					Password: "pass",
				},
			},
			offset:      0,
			limit:       5,
			validResult: []int{0, 1},
		},

		{
			name: "offset grater users",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
			},
			offset:      3,
			limit:       5,
			validResult: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			all, _ := user.NewService(user2.NewStorage(&tc.mp)).GetAll(context.TODO(), tc.limit, tc.offset)
			assert.Equal(t, len(tc.validResult), len(all))
		})
	}
}

func TestUser_CreateUser(t *testing.T) {
	testCases := []struct {
		name    string
		mp      map[int]*user.User
		u       func() *user.User
		isValid bool
	}{
		{
			name: "valid user",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       1,
					Name:     "BLA",
					Password: "pass",
				},
			},
			u: func() *user.User {
				return user3.TestValidUser(t)
			},
			isValid: true,
		},

		{
			name: "empty user",
			mp:   map[int]*user.User{},
			u: func() *user.User {
				return user3.TestEmptyUser(t)
			},
			isValid: true, //must be false
		},

		{
			name: "invalid user",
			mp:   map[int]*user.User{},
			u: func() *user.User {
				return user3.TestInvalidUser(t)
			},
			isValid: true, //must be false
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := tc.u()
			dto := &user.CreateUserDto{
				Name:     u.Name,
				Email:    u.Email,
				Password: u.Password,
			}
			err := user.NewService(user2.NewStorage(&tc.mp)).Create(context.TODO(), dto)
			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUser_Delete(t *testing.T) {
	testCases := []struct {
		name    string
		mp      map[int]*user.User
		uid     int
		isValid bool
	}{
		{
			name: "valid user",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       1,
					Name:     "BLA",
					Password: "pass",
				},
			},
			uid:     0,
			isValid: true,
		},

		{
			name:    "empty users",
			mp:      map[int]*user.User{},
			uid:     1,
			isValid: false,
		},

		{
			name: "user does not exist",
			mp: map[int]*user.User{
				0: {
					Id:       0,
					Name:     "BLA",
					Password: "pass",
				},
				1: {
					Id:       1,
					Name:     "BLA",
					Password: "pass",
				},
			},
			uid:     2,
			isValid: false, //must be false
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := user.NewService(user2.NewStorage(&tc.mp)).Delete(context.TODO(), tc.uid)
			if tc.isValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestUser_EncryptPassword(t *testing.T) {
	testCases := []struct {
		name string
		u    func() *user.User
	}{
		{
			name: "valid user",
			u: func() *user.User {
				return user3.TestValidUser(t)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := tc.u()
			err := u.EncryptPassword()
			assert.NoError(t, err)
			assert.Equal(t, u.Password, "")
			assert.NotEqual(t, u.EncryptedPassword, "")
		})
	}
}
