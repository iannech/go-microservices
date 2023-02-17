package handlers

import (
	"encoding/json"
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
	productList := models.GetProducts()

	data, err := json.Marshal(productList)
	if err != nil {
		http.Error(w, "Unable to marshall productList", http.StatusInternalServerError)
	} 

	w.Write(data)
}