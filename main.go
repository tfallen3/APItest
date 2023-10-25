package main

import (
	"ProductApi/measure"
	"ProductApi/product"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	productHandler := product.NewHandler()

	measureHandler := measure.NewHandler()

	r.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	r.HandleFunc("/product/{id}", productHandler.GetProduct).Methods("GET")
	r.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/measures", measureHandler.GetMeasures).Methods("GET")
	r.HandleFunc("/measure/{id}", measureHandler.GetMeasure).Methods("GET")
	r.HandleFunc("/measures", measureHandler.CreateMeasure).Methods("POST")
	r.HandleFunc("/measures/{id}", measureHandler.UpdateMeasure).Methods("PUT")
	r.HandleFunc("/measures/{id}", measureHandler.DeleteMeasure).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
