package main

import (
	"fmt"
	"sort"
)

func Joltages() {
	data := ReadData("day_10.txt")
	joltages := arrayAtoi(data)

	sort.Ints(joltages)

	totalJoltage := 0
	joltDifference1 := 0
	// Starts at 1 because the Phone is 3 jolts higher than the highest entry
	joltDifference3 := 1

	for _, joltage := range joltages {
		joltDifference := joltage - totalJoltage

		switch joltDifference {
		case 1:
			joltDifference1++
		case 3:
			joltDifference3++
		}

		totalJoltage = joltage
	}

	fmt.Printf("1-jolt differences: %v, 3-jolt differences: %v, multiplied: %v", joltDifference1, joltDifference3, joltDifference1*joltDifference3)
}
