package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

type Inventory struct {
	ID          int    `json:"id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
}

var conn *pgx.Conn

func getPort() int {
	portStr := os.Getenv("PORT")
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			return port
		}
	}
	return 8080 // Default port
}

func getDbURI() string {
	// Get database connection parameters from environment variables
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	// Construct the database connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	return dbURI
}

func main() {
	// Connect to the database
	var err error
	dbUri := getDbURI()
	// TODO: DB connection is attempted once. If DB is not ready or goes down later, it should recover
	conn, err = pgx.Connect(context.Background(), dbUri)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	r := mux.NewRouter()

	r.HandleFunc("/statusz", handleStatusz)
	r.HandleFunc("/inventory", handleInventory)
	port := getPort()
	log.Printf("Starting http server at port %d", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}

func handleStatusz(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.RequestURI)
	response := map[string]string{
		"message": "hello world",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleInventory(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request at %s", r.Method, r.RequestURI)
	switch r.Method {
	case http.MethodGet:
		listInventory(w, r)
	case http.MethodPost:
		createInventory(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func listInventory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Query the database
	rows, err := conn.Query(context.Background(), "SELECT id, product_name, quantity FROM inventory")
	if err != nil {
		http.Error(w, "Error querying database", http.StatusInternalServerError)
		log.Printf("Error querying database: %v", err)
		return
	}
	defer rows.Close()

	// Parse the results
	var inventoryList []Inventory
	for rows.Next() {
		var item Inventory
		err := rows.Scan(&item.ID, &item.ProductName, &item.Quantity)
		if err != nil {
			http.Error(w, "Error parsing database result", http.StatusInternalServerError)
			log.Printf("Error parsing database result: %v", err)
			return
		}
		inventoryList = append(inventoryList, item)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over database results", http.StatusInternalServerError)
		log.Printf("Error iterating over database results: %v", err)
		return
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(inventoryList)
}

func createInventory(w http.ResponseWriter, r *http.Request) {
	var item Inventory
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert the new item into the database
	var id int
	err = conn.QueryRow(context.Background(),
		"INSERT INTO inventory (product_name, quantity) VALUES ($1, $2) RETURNING id",
		item.ProductName, item.Quantity).Scan(&id)
	if err != nil {
		http.Error(w, "Error creating inventory item", http.StatusInternalServerError)
		log.Printf("Error creating inventory item: %v", err)
		return
	}

	// Set the ID of the created item
	item.ID = id

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}
