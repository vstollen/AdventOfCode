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