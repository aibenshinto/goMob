package crud

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func CreateMobilePhone(w http.ResponseWriter, r *http.Request, db *sql.DB) {

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

	_, err = db.Exec("INSERT INTO mobilePhones (name, image, specs, price) VALUES ($1, $2, $3, $4)",
		name, imagePath, specs, price)
	if err != nil {
		http.Error(w, "Failed to create mobile phone", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "New mobile phone created: %s", name)
	w.WriteHeader(http.StatusCreated)
}
