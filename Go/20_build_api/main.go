package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Category struct {
	Name    string `json:"name"`
	TaxRate int    `json:"taxRate"`
}

type Item struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Price    float64  `json:"price"`
	Quantity int      `json:"quantity"`
}

var inventory []Item
var nextItemID int = 1

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/items", addItemToInventory).Methods("POST")
	r.HandleFunc("/items", viewInventory).Methods("GET")
	r.HandleFunc("/items/{id}", updateItemById).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItemById).Methods("DELETE")

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", r)
}

func addItemToInventory(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newItem.ID = nextItemID
	nextItemID++
	inventory = append(inventory, newItem)
	json.NewEncoder(w).Encode(newItem)
}

func viewInventory(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(inventory)
}

func updateItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var updatedItem Item
	err = json.NewDecoder(r.Body).Decode(&updatedItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, item := range inventory {
		if item.ID == id {
			inventory[i] = updatedItem
			inventory[i].ID = id // Ensure the ID remains unchanged
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}

func deleteItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	for i, item := range inventory {
		if item.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Fprintf(w, "Deleted item with ID: %d", id)
			return
		}
	}

	http.Error(w, "Item not found", http.StatusNotFound)
}
