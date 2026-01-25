package product

import (
	"golang-practice/internal/base/1-13-packages/common"
	"log"
	"time"
)

type Product struct {
	ID    int64
	Name  string
	Price float64
	Stock int
}

type ProductService struct {
	products []*Product
	logger   *log.Logger
}

var TestMessage string

func init() {
	TestMessage = "Hello World"
}

func NewProductService() *ProductService {
	return &ProductService{
		products: []*Product{},
		logger:   common.NewLogger("product"),
	}
}

func (ps *ProductService) AddProduct(name string, price float64, stock int) {
	product := &Product{
		ID:    time.Now().Unix(),
		Name:  name,
		Price: price,
		Stock: stock,
	}

	ps.products = append(ps.products, product)
	ps.logger.Printf("Added product %s\n", product.Name)
}

func (ps *ProductService) FindProductByID(id int64) *Product {
	for _, product := range ps.products {
		if product.ID == id {
			return product
		}
	}
	return nil
}

func (ps *ProductService) ListAllProducts() {
	for _, product := range ps.products {
		ps.logger.Printf("Listing product %s\n", product.Name)
	}
}
