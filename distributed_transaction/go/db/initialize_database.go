package database

import (
	"database/sql"
	"fmt"
	"os"
)

func initializeDB(db *sql.DB, filename string) error {
	// Read the SQL file
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %v", err)
	}

	// Execute the SQL commands in the file
	_, err = db.Exec(string(fileContent))
	if err != nil {
		return fmt.Errorf("error executing SQL file: %v", err)
	}

	return nil
}
