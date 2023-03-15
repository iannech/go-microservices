// Package classification of Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go-microservices/models"

	"github.com/gorilla/mux"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// A list of products returned in the response
// swagger: response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []models.Product
}

// Products handler for getting and updating products
type Products struct {
	l *log.Logger
	v *models.Validation
}

// NewProducts returns a new products handler with the given logger
func NewProducts(l *log.Logger, v *models.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func getProductID(r *http.Request) int {
	// parse the product id from url
	vars := mux.Vars(r)

	// convert id into integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
