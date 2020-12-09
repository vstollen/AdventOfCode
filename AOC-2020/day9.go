package main

import "fmt"

func XMASCypher() {
	data := ReadData("day_9.txt")

	numbers := arrayAtoi(data)

	invalidNumber := findInvalidNumber(numbers)
	fmt.Printf("No sum for %v\n", invalidNumber)

	setWithSum := findSetWithSum(invalidNumber, numbers)
	fmt.Printf("The smallest and largest number adding up to %v summed together is: %v", invalidNumber, setWithSum[0]+setWithSum[len(setWithSum)-1])
}

func findInvalidNumber(numbers []int) int {
	for i := 25; i < len(numbers); i++ {
		if !sumExists(numbers[i], numbers[i-25:i]) {
			return numbers[i]
		}
	}

	return -1
}

func findSetWithSum(number int, values []int) []int {
	for i, _ := range values {
		sum := values[i]

		if sum >= number {
			continue
		}

		for j := i + 1; j < len(values); j++ {
			sum += values[j]

			if sum > number {
				break
			}

			if sum == number {
				return values[i : j+1]
			}
		}
	}

	return []int{}
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

			if val1+val2 != number {
				continue
			}

			return true
		}
	}

	return false
}
