package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to AdventOfCode 2020!\n" +
		"Please select a day (1-6):")

	var day int

	_, _ = fmt.Scanln(&day)

	switch day {
	case 1:
		ExpenseReport()
	case 2:
		CountValidPasswords()
	case 3:
		CountTrees()
	case 4:
		CountValidPassports()
	case 5:
		FindHighestSeatId()
	case 6:
		CountYesAnsweredQuestions()
	}
}
