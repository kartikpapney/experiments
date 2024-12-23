package main

import (
	"airline/mod/controller"
	database "airline/mod/db"
	"airline/mod/utils"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

/*
	1. ✅ Have a MySQL db running on a docker instance. Expose the UI to some endpoint.
	2. ✅ Create a database named airline
	3. ✅ Implement the tables in MySQL DB.
	4. Initialize a way to dummy the database during the start of the project.
		- Can test for the below cases
			- 10 airlines
			- 10,00,000 users
			- 100 flights per airline
			- 120 seats per flight
			- (100 flight/ airline) * (10 airline) * 1 trip schedule
		- ✅ We'll mimic the behavior with 1 trip for now.
			- 1 airline
			- 120 users
			- 1 flight
			- 1 trip
			- 120 seats (120 goroutines try to get a random seat)
	5. ✅ Try to have a visualization on terminal for Seats after and before the booking. Also, the logs.
	6. ✅ Implement a random function to mimic multiple users booking seats behavior
	7. ✅ Implement the below cases
		- What happens without lock
		- What happens without skip lock
		- What happens with skip lock
	8. Doc the performance improvement on With/ Without Skip Lock
*/

func execute(db *sql.DB, users []*controller.User, command string) time.Duration {
	startTime := time.Now()
	if command == "normal" {
		for _, user := range users {
			controller.BookSeat(db, user)
		}
	} else if command == "lock" {
		var wg sync.WaitGroup
		for _, user := range users {
			wg.Add(1)
			go func() {
				defer wg.Done()
				controller.BookSeat(db, user)
			}()
		}
		wg.Wait()
	} else if command == "skip-lock" {
		var wg sync.WaitGroup
		for _, user := range users {
			wg.Add(1)
			go func() {
				defer wg.Done()
				controller.BookSeatWithSkipLock(db, user)
			}()
		}
		wg.Wait()
	}
	elapsedTime := time.Since(startTime)
	return elapsedTime
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("invalid arguments")
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting current directory: ", err)
	}

	db, err := database.Connect(fmt.Sprintf("%s/db/schema.sql", dir))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = database.InsertDummyData(db)
	if err != nil {
		log.Fatal(err)
	}

	users, err := controller.GetUser(db)
	if err != nil {
		log.Fatal(err)
	}

	duration := execute(db, users, os.Args[1])
	fmt.Printf("Time: %s\n", duration)

	seats, err := controller.GetSeats(db)
	utils.PrintSeats(seats)
	if err != nil {
		log.Fatal(err)
	}

}
