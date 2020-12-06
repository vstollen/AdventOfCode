package main

import "fmt"

type group []string

func CountYesAnsweredQuestions() {
	data := ReadData("day_6.txt")
	groups := parseGroups(data)

	sumAnyoneYes := 0
	sumEveryoneYes := 0

	for _, group := range groups {
		answered := map[rune]int{}

		for _, line := range group {
			for _, answer := range line {
				answered[answer]++
			}
		}

		for _, answer := range answered {
			if answer == len(group) {
				sumEveryoneYes++
			}
		}

		sumAnyoneYes += len(answered)
	}

	fmt.Printf("Sum of questions answered by anyone: %v\n", sumAnyoneYes)
	fmt.Printf("Sum of questions answered by everyone: %v\n", sumEveryoneYes)

}

func parseGroups(data []string) []group {
	var groups []group

	for _, line := range data {
		if len(groups) == 0 || line == "" {
			groups = append(groups, group{})

			if line == "" {
				continue
			}
		}

		groups[len(groups)-1] = append(groups[len(groups)-1], line)
	}

	return groups
}
