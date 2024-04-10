package crud

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateMobilePhone(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid mobile phone ID", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	specs := r.FormValue("specs")
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to get image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imagePath := "../images/" + handler.Filename
	outputFile, err := os.Create(imagePath)
	if err != nil {
		fmt.Println("error creating path", err)
		http.Error(w, "Failed to save image file", http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, file)
	if err != nil {
		http.Error(w, "Failed to write image file", http.StatusInternalServerError)
		return
	}

	// Set ID from URL parameter

	_, err = db.Exec("UPDATE mobilePhones SET name = $1, image = $2, specs = $3, price = $4 WHERE id = $5",
		name, imagePath, specs, price, id)
	if err != nil {
		http.Error(w, "Failed to update mobile phone", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
