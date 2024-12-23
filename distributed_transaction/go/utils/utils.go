package utils

import (
	"airline/mod/controller"
	"fmt"
)

func PrintSeats(seats []*controller.Seat) error {
	fmt.Println("---------------------------------------")
	for _, seat := range seats {
		if seat.UserId.Valid {
			fmt.Print("x ")
		} else {
			fmt.Print(". ")
		}

		if seat.Id%20 == 0 {
			fmt.Println()
		}
		if seat.Id == 60 {
			fmt.Println()
		}
	}
	fmt.Println("---------------------------------------")
	return nil
}
