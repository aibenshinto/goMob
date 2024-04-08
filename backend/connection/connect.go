package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// Function to establish a connection to the PostgreSQL database
func ConnectDB() (*sql.DB, error) {
	// PostgreSQL connection string
	connectionString := "postgres://vefceiis:pzTKYphme3DcRyDg3ertxItmrtc44HiQ@rain.db.elephantsql.com/vefceiis"

	// Open a connection to the database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to the database!")

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS mobilePhones (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		image TEXT,
		specs TEXT,
		price NUMERIC(10, 2)
	);
`

	// Execute the SQL statement to create the table
	db.Exec(createTableSQL)

	return db, nil
}
