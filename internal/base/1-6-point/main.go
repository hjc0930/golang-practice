package main

import "fmt"

type User struct {
	ID      int
	Name    string
	Balance float64
}

func main() {
	//user := User{ID: 1, Name: "Alice", Balance: 100}
	//
	//updateBalanceValue(user, 50)
	//fmt.Println(user.Balance)
	//
	//updateBalancePointer(&user, 100)
	//fmt.Println(user.Balance)

	//users := []*User{
	//	&User{1, "Alice", 100},
	//	&User{2, "Bob", 150},
	//	&User{3, "Charlie", 200},
	//}
	//
	//for _, user := range users {
	//	user.Balance += 50.0
	//}
	//
	//for _, user := range users {
	//	fmt.Println(user.Balance)
	//}

	/** *********************** 二级指针 *********************/
	a := 10
	p := &a
	pp := &p

	fmt.Println(a, *p, **pp)
	/** *********************** 二级指针 *********************/
}

func updateBalanceValue(user User, amount float64) {
	user.Balance += amount
}

func updateBalancePointer(user *User, amount float64) {
	user.Balance += amount
}
