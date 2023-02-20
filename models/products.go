package models

import (
	"encoding/json"
	"io"
	"time"
)

// model defining structure for an API product
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

func AddProduct(p *Product){
	p.ID = GetNextId()
	productList = append(productList, p)
}

func GetNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// Products is a collection of Product
type Products []*Product

// using Encoder() is much faster and better compared to Marshall
// as it doesn't have to buffer the output into an in memory slice
// of bytes. This reduces allocations and the overheads of the service
func(p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func(p *Product) FromJSON(r io.Reader) error{
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
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