package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"go-microservices/models"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GET: all Products
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList := models.GetProducts()

	// call custom method that uses Encoder() to return ProductList
	err := productList.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshall productList", http.StatusInternalServerError)
	}
}

// POST: add Product
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	// retrieve product payload from context
	product := r.Context().Value(KeyProduct{}).(models.Product)

	models.AddProduct(&product)
}

// PUT: Update Product
func (p Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Unable to convert id.", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT products ")

	// retrieve product payload from context
	product := r.Context().Value(KeyProduct{}).(models.Product)

	err = models.UpdatePoduct(id, &product)

	if err == models.ErrorProductNotFound {
		http.Error(w, "Product not found!", http.StatusNotFound)
	}

	if err != nil {
		http.Error(w, "Product not found!", http.StatusInternalServerError)
	}
}

// Context is used with keys
type KeyProduct struct{}

// Middleware validation of product
func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := models.Product{}

		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
