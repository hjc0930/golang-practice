package user

import (
	"fmt"
	"golang-practice/internal/base/1-13-packages/common"
	"time"
)

type User struct {
	ID       int64
	Username string
	Password string
	Email    string
}

func CreateUser(name, email, password string) *User {
	common.LogOperation("CreateUser")

	user := &User{
		ID:       time.Now().Unix(),
		Username: name,
		Password: password,
		Email:    email,
	}
	if !user.valiodatePassword(password) {
		fmt.Println("password error")
		return nil
	}

	return user
}

func (u *User) DisplayInfo() {
	fmt.Printf("ID:%d User:%s \n", u.ID, u.Username)
}

func (u *User) valiodatePassword(password string) bool {
	return len(u.Password) >= 6
}
