package crud

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"main.go/models"
)

func SearchMobilePhones(w http.ResponseWriter, r *http.Request, db *sql.DB) {
    queryParams := r.URL.Query()
    searchTerm := strings.TrimSpace(queryParams.Get("search"))

    if searchTerm == "" {
        http.Error(w, "Search term is required", http.StatusBadRequest)
        return
    }

    query := `
        SELECT id, name, image, specs, price
        FROM mobilePhones
        WHERE LOWER(name) LIKE LOWER($1)
    `

    rows, err := db.Query(query, "%"+searchTerm+"%")
    if err != nil {
        http.Error(w, "Failed to search for mobile phones", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var products []models.MobilePhone
    for rows.Next() {
        var phone models.MobilePhone
        err := rows.Scan(&phone.ID, &phone.Name, &phone.Image, &phone.Specs, &phone.Price)
        if err != nil {
            http.Error(w, "Failed to scan mobile phone row", http.StatusInternalServerError)
            return
        }
        products = append(products, phone)
    }

    err = rows.Err()
    if err != nil {
        http.Error(w, "Error iterating over search results", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    err = json.NewEncoder(w).Encode(products)
    if err != nil {
        http.Error(w, "Failed to encode products to JSON", http.StatusInternalServerError)
        return
    }
}
