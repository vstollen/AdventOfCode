package main

import "fmt"

func XMASCypher() {
	data := ReadData("day_9.txt")

	numbers := arrayAtoi(data)

	for i := 25; i < len(numbers); i++ {
		if !sumExists(numbers[i], numbers[i-25:i]) {
			fmt.Printf("No sum for %v\n", numbers[i])
			break
		}
	}
}

func sumExists(number int, values []int) bool {
	for _, val1 := range values {
		if val1 >= number {
			continue
		}

		for _, val2 := range values {
			if val1 == val2 {
				continue
			}

			if val1 + val2 != number {
				continue
			}

			return true
		}
	}

	return false
}
