package handlers

import (
	"log"
	"net/http"

	"go-microservices/models"
)

type Products struct{
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products{
	return &Products{l}
}

func(p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// handle request for a list of products
	if r.Method == http.MethodGet{
		p.getProducts(w, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// GET: all Products
func(p *Products) getProducts(w http.ResponseWriter, r *http.Request){
	productList := models.GetProducts()

	// call custom method that uses Encoder() to return ProductList 
	err := productList.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to marshall productList", http.StatusInternalServerError)
	} 
}

// POST: add Product
func(p *Products) addProduct(w http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST products")

	product := &models.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
	}

	models.AddProduct(product)
}