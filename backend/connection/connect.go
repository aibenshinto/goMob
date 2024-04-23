package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)


func ConnectDB() (*sql.DB, error) {
	// PostgreSQL connection string
	connectionString := "host=localhost port=5432 dbname=postgres user=postgres password=shinto36 connect_timeout=10 sslmode=disable"

	// Open a connection to the database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}



	fmt.Println("Connected to the database!")


// 	insertinto := `INSERT INTO users (username, password, role) VALUES ('user', 'user123', 'user');
// 	`
// 	db.Exec(insertinto)
// 	// Execute the SQL statement to create the table
// 	db.Exec(createTableSQL)

	return db, nil
}
