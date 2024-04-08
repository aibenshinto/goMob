package crud

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteMobilePhone(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid mobile phone ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("DELETE FROM mobilePhones WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to delete mobile phone", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

