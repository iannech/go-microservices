package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	// handle request to add product
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}

	// handle request to partially update product
	// Native Go isn't good at parsing URI
	if r.Method == http.MethodPut {
		// expect the id in the URI
		// use regex to parse
		reg := regexp.MustCompile(`/([0-9]+)`)
		group := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, err := strconv.Atoi(idString)
		
		if err != nil {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, w, r)
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

// PUT: Update Product
func(p Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT products")

	product := &models.Product{}

	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Unable to decode json", http.StatusBadRequest)
	}

	err = models.UpdatePoduct(id, product)

	if err == models.ErrorProductNotFound {
		http.Error(w, "Product not found!", http.StatusNotFound)
	}

	if err != nil {
		http.Error(w, "Product not found!", http.StatusInternalServerError)
	}
}