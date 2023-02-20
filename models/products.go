package models

import (
	"encoding/json"
	"io"
	"time"
)

// model representing a product
type Product struct{
	ID 				int 	`json:"id"`
	Name 			string	`json:"name"`
	Description		string	`json:"description"`
	Price 			float32	`json:"price"`
	SKU 			string	`json:"sku"`
	CreatedOn		string	`json:"-"` // do not output this field
	UpdatedOn		string	`json:"-"`
	DeletedOn		string	`json:"-"`
}

// returns all products
func GetProducts() Products{
	return productList
}

// using Encoder() is much faster and better compared to Marshall
// cleaner way to return products/error
type Products []*Product

func(p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

// static list of products to act as data source
var productList = [] *Product{
	&Product{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		Price: 2.45,
		SKU: "abc323",
		CreatedOn: time.Now().String(),
		UpdatedOn: time.Now().String(),
	},
	&Product{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		Price: 1.99,
		SKU: "fjd343",
		CreatedOn: time.Now().String(),
		UpdatedOn: time.Now().String(),
	},
}