package main

import (
	"fmt"
)

/**
 * + - * / && || < > <= >= == ! & | << >>
 */
func main() {
	price := 299

	discount := 0.8

	orderAmount := float64(price) * discount

	fmt.Println(orderAmount)
}
