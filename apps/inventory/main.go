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
	"github.com/sangharsh/devx-pipeline/inventory/db"
	"github.com/sangharsh/devx-pipeline/inventory/inventory"
)

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
	// TODO: DB connection is attempted once. If DB is not ready or goes down later, it should recover

	err := db.InitDB(context.Background(), getDbURI())
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.CloseDB()

	r := mux.NewRouter()

	r.HandleFunc("/statusz", handleStatusz)
	r.HandleFunc("/inventory", inventory.HandleInventory)
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
