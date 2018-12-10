package service

import (
	"encoding/json"
	"net/http"
	"stock/db"
	"stock/schema"

	"github.com/gorilla/mux"
)

func CreateStock(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var stock schema.Stock

	err := decoder.Decode(&stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = db.CreateStock(r.Context(), &stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func ModifyStock(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var item struct {
		ProductID string `json:"product_id" bson:"product_id"`
		Quantity  int    `json:"quantity" bson:"quantity"`
	}

	err := decoder.Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = db.ModifyStock(r.Context(), item.ProductID, item.Quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetStock(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	stock, err := db.GetStock(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stock)
}
