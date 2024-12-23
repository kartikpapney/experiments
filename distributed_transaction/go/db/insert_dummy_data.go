package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/go-faker/faker/v4"
)

// Insert dummy data into the database
func InsertDummyData(db *sql.DB) error {
	// Insert 1 airline
	_, err := db.Exec("INSERT INTO airlines (id, name) VALUES (1, 'Delta Airlines')")
	if err != nil {
		return fmt.Errorf("error inserting dummy airline: %v", err)
	}
	var wg sync.WaitGroup
	// Insert 120 users
	for i := 1; i <= 120; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err = db.Exec("INSERT INTO users (id, name) VALUES ($1, $2)", i, faker.Name())
		}()
	}
	wg.Wait()
	// Insert 1 flight
	_, err = db.Exec("INSERT INTO flights (id, airline_id, name) VALUES (1, 1, 'DL101')")
	if err != nil {
		return fmt.Errorf("error inserting dummy flight: %v", err)
	}

	// Insert 1 trip
	_, err = db.Exec("INSERT INTO trips (id, flight_id, start_time, end_time) VALUES (1, 1, '2024-12-24 08:00:00', '2024-12-24 12:00:00')")
	if err != nil {
		return fmt.Errorf("error inserting dummy trip: %v", err)
	}

	// Insert 120 seats
	for i := 1; i <= 120; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			seatLabel := fmt.Sprintf("%dA", i) // Example seat labels: "1A", "2A", etc.
			_, err = db.Exec("INSERT INTO seats (id, name, trip_id, user_id) VALUES ($1, $2, 1, $3)", i, seatLabel, nil)
		}()

	}
	wg.Wait()
	return nil
}
