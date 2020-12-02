package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CountValidPasswords() {
	data := ReadData("day_2.txt")

	validPasswords := 0
	for _, elem := range data {
		if isValid(elem) {
			validPasswords++
		}
	}

	fmt.Printf("There are %v valid passwords.", validPasswords)
}

func isValid(passwordWithPolicy string) bool {
	parts := strings.Split(passwordWithPolicy, " ")

	minmax := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(minmax[0])
	max, _ := strconv.Atoi(minmax[1])

	letter := rune(parts[1][0])
	password := parts[2]

	letterCount := 0
	for _, passwordLetter := range password {
		if passwordLetter == letter {
			letterCount++
		}
	}

	if letterCount < min {
		return false
	}

	if letterCount > max {
		return false
	}

	return true
}
