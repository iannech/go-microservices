package handlers

import (
	"go-microservices/models"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	// retrieve product payload from context
	product := r.Context().Value(KeyProduct{}).(models.Product)

	models.AddProduct(&product)
}
