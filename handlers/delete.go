package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ywpark1/microservice-in-go/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product
// responses:
//  201: noContentResponse

// DeleteProduct deletes a product from the database
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
