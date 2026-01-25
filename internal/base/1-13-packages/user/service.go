package user

import "golang-practice/internal/base/1-13-packages/common"

type UserService struct {
	users []*User
}

func NewUserService() *UserService {
	return &UserService{
		users: make([]*User, 0),
	}
}

func (us *UserService) AddUser(user *User) {
	if user != nil {
		us.users = append(us.users, user)
		common.LogOperation("Add user")
	}
}

func (us *UserService) ListAllUsers() {
	for _, user := range us.users {
		user.DisplayInfo()
	}
}
