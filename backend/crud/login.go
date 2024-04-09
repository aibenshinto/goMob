package crud

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"main.go/models"
)

func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {

		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return

	}

	query := "SELECT id, role FROM users WHERE username = $1 AND password = $2"

	var dbUser models.User
	err = db.QueryRow(query, user.Username, user.Password).Scan(&dbUser.ID, &dbUser.Role)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	switch dbUser.Role {
	case "admin":
		response := map[string]string{"role": "admin"}
		w.Header().Set("Content-Type", "application/json") // If user is admin, redirect to admin dashboard or send appropriate response
		json.NewEncoder(w).Encode(response)                // Example: http.Redirect(w, r, "/admin.html", http.StatusSeeOther)

		// Placeholder response for admin
	case "user":
		// If user is regular user, send success response with user role
		response := map[string]string{"role": "user"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	default:
		http.Error(w, "Unauthorized access", http.StatusUnauthorized)
	}

}
