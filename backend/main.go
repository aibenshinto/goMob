package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"main.go/connection"
	"main.go/crud"
)

func main() {
	// Connect to the database
	db, err := connection.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create a new router
	r := mux.NewRouter()

	// Register routes with handlers and pass db connection

	r.HandleFunc("/api/mobilephones", func(w http.ResponseWriter, r *http.Request) {
		crud.GetAllMobilePhones(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/mobilephones", func(w http.ResponseWriter, r *http.Request) {
		crud.CreateMobilePhone(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/api/mobilephones/{id}", func(w http.ResponseWriter, r *http.Request) {
		crud.GetMobilePhoneByID(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/api/mobilephones/{id}", func(w http.ResponseWriter, r *http.Request) {
		crud.UpdateMobilePhone(w, r, db)
	}).Methods("PUT")

	r.HandleFunc("/api/mobilephones/{id}", func(w http.ResponseWriter, r *http.Request) {
		crud.DeleteMobilePhone(w, r, db)
	}).Methods("DELETE")

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}), // Allow requests from any origin
	)

	// Create a CORS-enabled handler with the router
	corsHandler := cors(r)

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
