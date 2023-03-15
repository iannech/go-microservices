package handlers

import (
	"go-microservices/models"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (p Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id ", id)

	err := models.DeleteProduct(id)

	if err == models.ErrorProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
