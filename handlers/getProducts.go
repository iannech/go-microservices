package handlers

import (
	"go-microservices/models"
	"net/http"
)

// swagger:route GET /products listProducts
// Returns a list of products
// responses:
// 200: productsResponse
func (p *Products) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	productList := models.GetProducts()

	// call custom method that uses Encoder() to return ProductList
	err := models.ToJSON(productList, w)

	if err != nil {
		http.Error(w, "Unable to marshall productList", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//
//	200: productResponse
//	404: errorResponse
func (p *Products) GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := models.GetProductByID(id)

	switch err {
	case nil:

	case models.ErrorProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = models.ToJSON(prod, w)
	if err != nil {
		// we should never be here but log the error just incase
		p.l.Println("[ERROR] serializing product", err)
	}
}
