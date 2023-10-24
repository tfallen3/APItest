package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	UnitCost string `json:"'unit_cost"`
	Measure  string `json:"'measure' Единица измерения"`
}

type Measure struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var products []Product
var measures []Measure

func getMeasures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(measures)
}
func getMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range measures {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func createMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var measure Measure
	_ = json.NewDecoder(r.Body).Decode(&measure)
	measure.ID = strconv.Itoa(rand.Intn(20))
	measures = append(measures, measure)
	json.NewEncoder(w).Encode(measure)
}
func updateMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var measure Measure
	for index, item := range measures {
		if item.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&measure)
			measures[index].Name = measure.Name
			break
		}
	}
	json.NewEncoder(w).Encode(measures)
}

func deleteMeasure(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range measures {
		if item.ID == params["id"] {
			measures = append(measures[:index], measures[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(measures)
}
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(1000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var product Product
	for index, item := range products {
		if item.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&product)
			products[index].Name = product.Name
			products[index].Quantity = product.Quantity
			products[index].UnitCost = product.UnitCost
			products[index].Measure = product.Measure
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}
func main() {
	r := mux.NewRouter()
	products = append(products, Product{ID: "1", Name: "Апельсин", Quantity: 2, UnitCost: "16", Measure: "Килограмм"})
	products = append(products, Product{ID: "2", Name: "Шоколад", Quantity: 2, UnitCost: "126", Measure: "Килограмм"})
	measures = append(measures, Measure{ID: "1", Name: "Килограмм"})
	measures = append(measures, Measure{ID: "2", Name: "Штука"})
	r.HandleFunc("/products", getProducts).Methods("GET")
	r.HandleFunc("/product/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", createProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	r.HandleFunc("/measures", getMeasures).Methods("GET")
	r.HandleFunc("/measure/{id}", getMeasure).Methods("GET")
	r.HandleFunc("/measures", createMeasure).Methods("POST")
	r.HandleFunc("/measures/{id}", updateMeasure).Methods("PUT")
	r.HandleFunc("/measures/{id}", deleteMeasure).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
