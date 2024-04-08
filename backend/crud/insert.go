package crud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"main.go/models"
)

func CreateMobilePhone(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	
	var phone models.MobilePhone

	err := json.NewDecoder(r.Body).Decode(&phone)
	
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO mobilePhones (name, image, specs, price) VALUES ($1, $2, $3, $4)",
		phone.Name, phone.Image, phone.Specs, phone.Price)
	if err != nil {
		http.Error(w, "Failed to create mobile phone", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w,"the data is  : %v",phone)

	w.WriteHeader(http.StatusCreated)
}
