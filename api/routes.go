package api

import (
	"net/http"

	"inventory-api/internal/product"

	"github.com/gorilla/mux"
)

func SetupRoutes(handler *product.Handler) http.Handler {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/products", handler.CreateProduct).Methods("POST")
	v1.HandleFunc("/products/{id}", handler.UpdateProduct).Methods("PUT")
	v1.HandleFunc("/products/{id}", handler.GetProduct).Methods("GET")
	v1.HandleFunc("/products/{id}", handler.DeleteProduct).Methods("DELETE")

	return r
}
