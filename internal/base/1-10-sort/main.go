package main

import (
	"fmt"
	"sort"
	"time"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Rating   float64
	Sales    int
	CreateAt time.Time
}

type ProductManager struct {
	Products []Product
}

func NewProductManager() *ProductManager {
	return &ProductManager{Products: []Product{
		{ID: 1, Name: "Product A", Price: 29.99, Rating: 4.5, Sales: 150, CreateAt: time.Date(2022, time.January, 10, 0, 0, 0, 0, time.UTC)},
		{ID: 2, Name: "Product B", Price: 19.99, Rating: 4.0, Sales: 200, CreateAt: time.Date(2022, time.February, 15, 0, 0, 0, 0, time.UTC)},
		{ID: 3, Name: "Product C", Price: 39.99, Rating: 4.8, Sales: 120, CreateAt: time.Date(2022, time.March, 20, 0, 0, 0, 0, time.UTC)},
		{ID: 4, Name: "Product D", Price: 24.99, Rating: 3.9, Sales: 300, CreateAt: time.Date(2022, time.April, 25, 0, 0, 0, 0, time.UTC)},
	}}
}

func (pm *ProductManager) sortByPrice(asced bool) {
	if asced {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].Price < pm.Products[j].Price
		})
	} else {
		sort.Slice(pm.Products, func(i, j int) bool {
			return pm.Products[i].Price > pm.Products[j].Price
		})
	}

}

func main() {
	pm := NewProductManager()
	pm.sortByPrice(false)
	for _, product := range pm.Products {
		fmt.Println(product.Price, product.Name, product.Price)
	}
}
