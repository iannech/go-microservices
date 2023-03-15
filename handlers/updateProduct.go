package handlers

import (
	"go-microservices/models"
	"net/http"
)

func (p Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	// retrieve product payload from context
	product := r.Context().Value(KeyProduct{}).(models.Product)
	p.l.Println("[DEBUG] updating record id ", product.ID)

	err := models.UpdatePoduct(product)

	if err == models.ErrorProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(&GenericError{Message: "Product not found in database"}, w)
		return
	}

	// write the no content success header
	w.WriteHeader(http.StatusNoContent)
}
