package main

import "fmt"

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	cartList := make([]CartItem, 0)

	cartList = append(cartList, CartItem{
		Name:     "Phone",
		Price:    699.99,
		Quantity: 10,
	}, CartItem{
		Name:     "Laptop",
		Price:    300.00,
		Quantity: 12,
	}, CartItem{
		Name:     "Mouse",
		Price:    25.50,
		Quantity: 15,
	})

	total := 0.0
	totalCard := 0
	for _, item := range cartList {
		total += item.Price * float64(item.Quantity)
		totalCard += item.Quantity
	}

	fmt.Printf("Total Price: %.2f\n", total)
	fmt.Println("Total Quantity: ", totalCard)

	// remove first item
	cartList = cartList[1:]
}
