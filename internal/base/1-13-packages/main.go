package main

import (
	"fmt"
	"golang-practice/internal/base/1-13-packages/product"
	"golang-practice/internal/base/1-13-packages/user"
)

func main() {
	userService := user.NewUserService()

	fmt.Println("用户服务测试")
	userService.AddUser(user.CreateUser("cheng", "cheng@qq.com", "123456"))
	userService.AddUser(user.CreateUser("cheng2", "cheng2@qq.com", "123456"))

	userService.ListAllUsers()

	productService := product.NewProductService()

	fmt.Println("商品服务测试")
	productService.AddProduct("iPhone 17", 6000.00, 10)
	productService.AddProduct("iPhone 18", 7000.00, 10)

	productService.ListAllProducts()

	fmt.Println(product.TestMessage)
}
