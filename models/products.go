package models

import (
	"fmt"
	"time"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

// model defining structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"` // do not output this field
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// returns all products from Database
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrorProductNotFound
	}

	return productList[i], nil
}

// AddProduct adds a new product to the database
func AddProduct(p Product) {
	maxID := productList[len(productList)-1].ID
	p.ID = maxID + 1
	productList = append(productList, &p)
}

// UpdateProduct replaces a product in the database with the given
// item.
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdatePoduct(p Product) error {
	i := findIndexByProductID(p.ID)

	if i == -1 {
		return ErrorProductNotFound
	}

	// update the product in the DB
	productList[i] = &p

	return nil
}

// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)

	if i == -1 {
		return ErrorProductNotFound
	}

	productList = append(productList[:1], productList[i+1])

	return nil
}

var ErrorProductNotFound = fmt.Errorf("Product not found!")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, 0, ErrorProductNotFound
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

// Products is a collection of Product
type Products []*Product

// static list of products to act as data source
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd343",
		CreatedOn:   time.Now().String(),
		UpdatedOn:   time.Now().String(),
	},
}
