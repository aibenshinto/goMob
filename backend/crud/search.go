package crud

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"main.go/models"
)

func GetMobilePhoneByID(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid mobile phone ID", http.StatusBadRequest)
		return
	}

	var phone models.MobilePhone
	err = db.QueryRow("SELECT id, name, image, specs, price FROM mobilePhones WHERE id = $1", id).
		Scan(&phone.ID, &phone.Name, &phone.Image, &phone.Specs, &phone.Price)
	if err != nil {
		http.Error(w, "Failed to get mobile phone", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(phone)
}
