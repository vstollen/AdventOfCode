package main

import (
	"fmt"
)

func FindHighestSeatId() {
	data := ReadData("day_5.txt")

	highestSeatId := -1
	for _, boardingPass := range data {
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
		if rowLow*8+seatLow > highestSeatId {
			highestSeatId = rowLow*8 + seatLow
		}
	}

	fmt.Printf("Highest Seat ID: %v\n", highestSeatId)

}
