package main

import "fmt"

type User struct {
	ID       int64
	UserName string
	Password string
	Email    string
	Age      int
	Address  Address
}

type Address struct {
	Province string
	City     string
	Street   string
	ZipCode  string
}

type UserProfile struct {
	ID       int
	UserName string
	Address
}

func main() {
	addr := Address{
		Province: "California",
		City:     "Los Angeles",
		Street:   "Sunset Blvd",
		ZipCode:  "90001",
	}

	user := User{
		ID:       1,
		UserName: "john_doe",
		Password: "securepassword",
		Email:    "john.doe@example.com",
		Age:      30,
		Address:  addr,
	}

	fmt.Print(user.Address.City)

}
