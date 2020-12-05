package main

import (
	"fmt"
)

func FindHighestSeatId() {
	data := ReadData("day_5.txt")

	occupiedSeats := map[int]bool{}

	highestSeatId := -1
	for _, boardingPass := range data {

		seatId := calcSeatId(boardingPass)

		occupiedSeats[seatId] = true

		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	fmt.Printf("Highest Seat ID: %v\n", highestSeatId)

	for i := 0; i < highestSeatId; i++ {
		if !occupiedSeats[i] {
			fmt.Printf("Seat ID %v is free.\n", i)
		}
	}
}

func calcSeatId(boardingPass string) int {
	rowLow := 0
	rowHigh := 127

	seatLow := 0
	seatHigh := 7

	for _, letter := range boardingPass {
		switch letter {
		case 'F':
			rowHigh = (rowLow + rowHigh) / 2
		case 'B':
			rowLow = ((rowLow + rowHigh) / 2) + 1
		case 'L':
			seatHigh = (seatLow + seatHigh) / 2
		case 'R':
			seatLow = ((seatLow + seatHigh) / 2) + 1
		}
	}

	return rowLow*8 + seatLow
}
