package crud

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"main.go/models"
)

func GetAllMobilePhones(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	rows, err := db.Query("SELECT id, name, image, specs, price FROM mobilePhones")
	if err != nil {
		http.Error(w, "Failed to fetch mobile phones", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Slice to hold mobile phones
	var phones []models.MobilePhone

	// Iterate over rows and scan into MobilePhone struct
	for rows.Next() {
		var phone models.MobilePhone
		err := rows.Scan(&phone.ID, &phone.Name, &phone.Image, &phone.Specs, &phone.Price)
		if err != nil {
			http.Error(w, "Failed to fetch mobile phones", http.StatusInternalServerError)
			return
		}
		phones = append(phones, phone)
	}

	// Check for any errors encountered during iteration
	err = rows.Err()
	if err != nil {
		http.Error(w, "Failed to fetch mobile phones", http.StatusInternalServerError)
		return
	}

	// Encode phones slice as JSON and send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(phones)

}
