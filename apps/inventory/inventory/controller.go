package inventory

import (
	"encoding/json"
	"log"
	"net/http"
)

func HandleInventory(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.RequestURI)
	switch r.Method {
	case http.MethodGet:
		ProcessListInventory(w, r)
	case http.MethodPost:
		ProcessCreateInventory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ProcessListInventory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	list, err := ListInventory()

	if err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		return
	}
	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

func ProcessCreateInventory(w http.ResponseWriter, r *http.Request) {
	var item Inventory
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	item, err = CreateInventory(item)
	if err != nil {
		http.Error(w, "Error creating inventory item", http.StatusInternalServerError)
		log.Printf("Error creating inventory item: %v", err)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
