package crud

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/models"
)

func UpdateMobilePhone(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid mobile phone ID", http.StatusBadRequest)
		return
	}

	var phone models.MobilePhone
	err = json.NewDecoder(r.Body).Decode(&phone)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	phone.ID = id // Set ID from URL parameter

	_, err = db.Exec("UPDATE mobilePhones SET name = $1, image = $2, specs = $3, price = $4 WHERE id = $5",
		phone.Name, phone.Image, phone.Specs, phone.Price, phone.ID)
	if err != nil {
		http.Error(w, "Failed to update mobile phone", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
