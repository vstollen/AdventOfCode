package main

import (
	"fmt"
	"strconv"
	"strings"
)

type password struct {
	pos1, pos2 int
	letter     rune
	password   string
}

func CountValidPasswords() {
	stringData := ReadData("day_2.txt")
	data := parsePasswords(stringData)

	validPasswordsOld := 0
	validPasswordsNew := 0
	for _, elem := range data {
		if elem.oldIsValid() {
			validPasswordsOld++
		}
		if elem.newIsValid() {
			validPasswordsNew++
		}
	}

	fmt.Printf("There are %v valid passwords according to the old policy.\n", validPasswordsOld)
	fmt.Printf("There are %v valid passwords according to the new policy.\n", validPasswordsNew)
}

// Checker for the old password policy
func (p password) oldIsValid() bool {

	letterCount := 0
	for _, letter := range p.password {
		if letter == p.letter {
			letterCount++
		}
	}

	if letterCount < p.pos1 {
		return false
	}

	if letterCount > p.pos2 {
		return false
	}

	return true
}

func (p password) newIsValid() bool {
	firstEqualsLetter := rune(p.password[p.pos1-1]) == p.letter
	secondEqualsLetter := rune(p.password[p.pos2-1]) == p.letter

	return firstEqualsLetter != secondEqualsLetter
}

func parsePasswords(data []string) []password {
	var passwords = make([]password, len(data))

	for i, _ := range data {

		parts := strings.Split(data[i], " ")

		numbers := strings.Split(parts[0], "-")
		pos1, _ := strconv.Atoi(numbers[0])
		pos2, _ := strconv.Atoi(numbers[1])

		letter := rune(parts[1][0])
		passwordText := parts[2]

		passwords[i] = password{
			pos1:     pos1,
			pos2:     pos2,
			letter:   letter,
			password: passwordText,
		}
	}

	return passwords
}
