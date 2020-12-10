package main

import (
	"fmt"
	"sort"
)

func Joltages() {
	data := ReadData("day_10.txt")
	joltages := arrayAtoi(data)

	sort.Ints(joltages)

	joltages = append(joltages, joltages[len(joltages)-1] + 3)

	countDifferences(joltages)

	distinctArrangements := countDistinctArrangements(joltages)

	fmt.Printf("Distinct Arrangements: %v", distinctArrangements)
}

// countDifferences requires the joltages to be sorted
func countDifferences(joltages []int) {
	totalJoltage := 0
	joltDifference1 := 0
	// Starts at 1 because the Phone is 3 jolts higher than the highest entry
	joltDifference3 := 0

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

	fmt.Printf("1-jolt differences: %v, 3-jolt differences: %v, multiplied: %v\n", joltDifference1, joltDifference3, joltDifference1*joltDifference3)
}

// countDistinctArrangements requires the joltages to be sorted
func countDistinctArrangements(joltages []int) int  {
	arrangementCounts := map[int]int{0: 1}

	joltages = append([]int{0}, joltages...)

	for i := 0; i < len(joltages); i++ {
		for j := i+1; j < len(joltages) && joltages[j] <= joltages[i]+3; j++ {
			arrangementCounts[j] += arrangementCounts[i]
		}
	}

	return arrangementCounts[len(joltages)-1]
}
