package inventory

import (
	"context"
	"log"

	"github.com/sangharsh/devx-pipeline/inventory/db"
)

func ListInventory() ([]Inventory, error) {
	// Query the database
	rows, err := db.GetPool().Query(context.Background(), "SELECT id, product_name, quantity FROM inventory")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return nil, err
		// http.Error(w, "Error querying database", http.StatusInternalServerError)
	}
	defer rows.Close()

	// Parse the results
	var list []Inventory
	for rows.Next() {
		var item Inventory
		err := rows.Scan(&item.ID, &item.ProductName, &item.Quantity)
		if err != nil {
			log.Printf("Error parsing database result: %v", err)
			return nil, err
			// http.Error(w, "Error parsing database result", http.StatusInternalServerError)
		}
		list = append(list, item)
	}
	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over database results: %v", err)
		return nil, err
		// http.Error(w, "Error iterating over database results", http.StatusInternalServerError)
	}
	return list, nil
}

func CreateInventory(item Inventory) (Inventory, error) {
	var id int
	// Insert the new item into the database
	err := db.GetPool().QueryRow(context.Background(),
		"INSERT INTO inventory (product_name, quantity) VALUES ($1, $2) RETURNING id",
		item.ProductName, item.Quantity).Scan(&id)

	if err != nil {
		return item, err
	}
	// Set the ID of the created item
	item.ID = id
	return item, err
}
