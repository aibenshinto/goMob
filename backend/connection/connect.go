package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)


func ConnectDB() (*sql.DB, error) {
	// PostgreSQL connection string
	connectionString := "postgres://vefceiis:pzTKYphme3DcRyDg3ertxItmrtc44HiQ@rain.db.elephantsql.com/vefceiis"

	// Open a connection to the database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Ping is used for checking wheather the connection is made or not
	// err = db.Ping()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to ping database: %v", err)
	// }

	fmt.Println("Connected to the database!")

// 	createTableSQL := `
// 	CREATE TABLE IF NOT EXISTS Users (
// 		id SERIAL PRIMARY KEY,
//     username VARCHAR(255) UNIQUE NOT NULL,
//     password VARCHAR(255) NOT NULL,
//     role VARCHAR(50) NOT NULL
// 	);
// `
// 	insertinto := `INSERT INTO users (username, password, role) VALUES ('user', 'user123', 'user');
// 	`
// 	db.Exec(insertinto)
// 	// Execute the SQL statement to create the table
// 	db.Exec(createTableSQL)

	return db, nil
}
