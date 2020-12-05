package main

import (
	"fmt"
	"strconv"
)

func ExpenseReport() {
	stringData := ReadData("day_1.txt")
	data := arrayAtoi(stringData)

	twoNumbers(data)
	threeNumbers(data)
}

func twoNumbers(data []int) {
	fmt.Println("Starting search for two numbers.")
	for _, elem1 := range data {
		for _, elem2 := range data {
			if elem1 == elem2 {
				continue
			}

			if elem1+elem2 == 2020 {
				solution := elem1 * elem2

				fmt.Printf("%v + %v = 2020\n", elem1, elem2)
				fmt.Printf("Solution: %v * %v = %v\n", elem1, elem2, solution)
				return
			}
		}
	}
}

func threeNumbers(data []int) {
	fmt.Printf("Starting search for three numbers.")

	for _, elem1 := range data {
		for _, elem2 := range data {
			if elem1 == elem2 {
				continue
			}

			if elem1+elem2 >= 2020 {
				continue
			}

			for _, elem3 := range data {
				if elem3 == elem1 || elem3 == elem2 {
					continue
				}

				if elem1+elem2+elem3 == 2020 {
					solution := elem1 * elem2 * elem3

					fmt.Printf("%v + %v + %v = 2020\n", elem1, elem2, elem3)
					fmt.Printf("Solution: %v * %v * %v = %v", elem1, elem2, elem3, solution)
					return
				}
			}

		}
	}
}

func arrayAtoi(data []string) []int {
	var convertedData = make([]int, len(data))

	for i, _ := range data {
		convertedData[i], _ = strconv.Atoi(data[i])
	}

	return convertedData
}
