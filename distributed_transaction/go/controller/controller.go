package controller

import (
	"database/sql"
	"fmt"
)

type Seat struct {
	Id     int
	Name   string
	TripId int
	UserId sql.NullInt64
}

type User struct {
	Id   int
	Name string
}

func GetSeats(db *sql.DB) ([]*Seat, error) {
	txn, err := db.Begin()
	if err != nil {
		txn.Rollback()
		return nil, fmt.Errorf("could not begin transaction: %v", err)
	}

	rows, err := txn.Query(`
		SELECT id, name, trip_id, user_id
		FROM seats
		WHERE trip_id = 1
		ORDER BY id;
	`)
	if err != nil {
		txn.Rollback()
		return nil, fmt.Errorf("could not query seats: %v", err)
	}
	var seats []*Seat

	for rows.Next() {
		var seat Seat
		err := rows.Scan(&seat.Id, &seat.Name, &seat.TripId, &seat.UserId)
		if err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		seats = append(seats, &seat)
	}

	if err := rows.Err(); err != nil {
		txn.Rollback()
		return nil, fmt.Errorf("error iterating over rows: %v", err)
	}

	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		return nil, fmt.Errorf("could not commit transaction: %v", err)
	}

	return seats, nil
}

func BookSeat(db *sql.DB, user *User) error {
	txn, _ := db.Begin()

	row := txn.QueryRow(`
		SELECT id, name, trip_id, user_id
		FROM seats
		WHERE trip_id = 1 AND user_id IS null
		ORDER BY id
		LIMIT 1
		FOR UPDATE
	`)

	var currSeat Seat
	err := row.Scan(&currSeat.Id, &currSeat.Name, &currSeat.TripId, &currSeat.UserId)

	if err != nil {
		txn.Rollback()
		return fmt.Errorf("could not read the seat: %v", err)
	}
	_, err = txn.Exec(`
		UPDATE seats
		SET user_id = $1
		WHERE id = $2
	`, user.Id, currSeat.Id)

	if err != nil {
		txn.Rollback()
		return fmt.Errorf("could not update the seat: %v", err)
	}
	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}

func BookSeatWithSkipLock(db *sql.DB, user *User) error {
	txn, _ := db.Begin()

	row := txn.QueryRow(`
		SELECT id, name, trip_id, user_id
		FROM seats
		WHERE trip_id = 1 AND user_id IS null
		ORDER BY id
		LIMIT 1
		FOR UPDATE
		SKIP LOCKED
	`)

	var currSeat Seat
	err := row.Scan(&currSeat.Id, &currSeat.Name, &currSeat.TripId, &currSeat.UserId)

	if err != nil {
		txn.Rollback()
		return fmt.Errorf("could not read the seat: %v", err)
	}
	_, err = txn.Exec(`
		UPDATE seats
		SET user_id = $1
		WHERE id = $2
	`, user.Id, currSeat.Id)

	if err != nil {
		txn.Rollback()
		return fmt.Errorf("could not update the seat: %v", err)
	}
	err = txn.Commit()
	if err != nil {
		txn.Rollback()
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}

func GetUser(db *sql.DB) ([]*User, error) {
	txn, _ := db.Begin()

	rows, _ := txn.Query(`
		SELECT id, name
		FROM users
	`)
	var users []*User
	for rows.Next() {
		var nUser User
		rows.Scan(&nUser.Id, &nUser.Name)
		users = append(users, &nUser)
	}

	err := txn.Commit()
	if err != nil {
		txn.Rollback()
		return nil, fmt.Errorf("could not commit transaction: %v", err)
	}
	return users, nil
}
