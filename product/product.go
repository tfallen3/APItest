package product

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	UnitCost string `json:"unit_cost"`
	Measure  string `json:"measure"`
}

type Handler struct {
	products []Product
}

func NewHandler() *Handler {
	ph := &Handler{}
	ph.products = append(ph.products, Product{ID: "1", Name: "Апельсин", Quantity: 2, UnitCost: "16", Measure: "Килограмм"})
	ph.products = append(ph.products, Product{ID: "2", Name: "Яблоко", Quantity: 2, UnitCost: "16", Measure: "Килограмм"})
	return ph
}

func (ph *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ph.products)
}
func (ph *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range ph.products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func (ph *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	// Найти максимальный существующий ID
	maxID := 0
	for _, existingProduct := range ph.products {
		productID, err := strconv.Atoi(existingProduct.ID)
		if err == nil && productID > maxID {
			maxID = productID
		}
	}

	product.ID = strconv.Itoa(maxID + 1)

	ph.products = append(ph.products, product)

	json.NewEncoder(w).Encode(product)
}
func (ph *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	params := mux.Vars(r)
	for index, item := range ph.products {
		if item.ID == params["id"] {
			_ = json.NewDecoder(r.Body).Decode(&product)
			ph.products[index].Name = product.Name
			ph.products[index].Quantity = product.Quantity
			ph.products[index].UnitCost = product.UnitCost
			ph.products[index].Measure = product.Measure
			break
		}
	}
	json.NewEncoder(w).Encode(ph.products)
}

func (ph *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range ph.products {
		if item.ID == params["id"] {
			ph.products = append(ph.products[:index], ph.products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(ph.products)
}
