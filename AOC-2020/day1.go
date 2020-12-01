package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ExpenseReport() {
	data := readData()

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

			if elem1 + elem2 == 2020 {
				solution := elem1 * elem2

				fmt.Printf("%v + %v = 2020\n", elem1, elem2)
				fmt.Printf("Solution: %v * %v = %v", elem1, elem2, solution)
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

			if elem1 + elem2 >= 2020 {
				continue
			}

			for _, elem3 := range data {
				if elem3 == elem1 || elem3 == elem2 {
					continue
				}

				if elem1 + elem2 + elem3 == 2020 {
					solution := elem1 * elem2 * elem3

					fmt.Printf("%v + %v + %v = 2020\n", elem1, elem2, elem3)
					fmt.Printf("Solution: %v * %v * %v = %v", elem1, elem2, elem3, solution)
					return
				}
			}

		}
	}
}

func readData() []int {

	file, err := os.Open("input/day_1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		
		data = append(data, number)
	}

	return data
}