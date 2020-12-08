package main

import (
	"fmt"
	"strconv"
	"strings"
)

type bag struct {
	name     string
	contains map[*bag]int
	heldBy   map[string]*bag
}

func CountOuterBagColors() {
	data := ReadData("day_7.txt")
	bagRules := parseBagRules(data)

	outerBags := map[*bag]bool{}

	outerBags = appendOuterParents(bagRules["shiny gold"], outerBags)
	delete(outerBags, bagRules["shiny gold"])

	for name, _ := range outerBags {
		fmt.Printf("%v\n", name.name)
	}

	fmt.Printf("Total: %v\n", len(outerBags))

	bagsInside := bagsNeededWithBag(bagRules["shiny gold"], bagRules)
	// Subtract the shiny gold bag to get the number of bags inside
	bagsInside--

	fmt.Printf("Bags needed inside Shiny Gold: %v", bagsInside)
}

func appendOuterParents(innerBag *bag, outerBags map[*bag]bool) map[*bag]bool {
	outerBags[innerBag] = true

	for _, outerBag := range innerBag.heldBy {
		outerBags = appendOuterParents(outerBag, outerBags)
	}

	return outerBags
}

func bagsNeededWithBag(outerBag *bag, rules map[string]*bag) int {
	sum := 1
	for innerBag, count := range outerBag.contains {
		sum += count * bagsNeededWithBag(innerBag, rules)
	}

	return sum
}

func (b bag) String() string {
	return b.name
}

func parseBagRules(data []string) map[string]*bag {
	bags := map[string]*bag{}

	for _, rule := range data {
		bagName, rule := splitAtFirst(rule, " bags contain ")
		outerBag, ok := bags[bagName]
		if !ok {
			bags[bagName] = &bag{name: bagName}
			bags[bagName].heldBy = make(map[string]*bag)
			bags[bagName].contains = make(map[*bag]int)
		}
		outerBag = bags[bagName]

		rule = strings.TrimSuffix(rule, ".")

		rules := strings.Split(rule, ", ")

		for _, rule := range rules {
			rule = strings.TrimSuffix(rule, " bag")
			rule = strings.TrimSuffix(rule, " bags")

			count, name := splitAtFirst(rule, " ")

			bagPointer, ok := bags[name]
			if !ok {
				bags[name] = &bag{
					name: name,
				}
				bagPointer = bags[name]
				bagPointer.heldBy = make(map[string]*bag)
				bagPointer.contains = make(map[*bag]int)
			}

			bagPointer.heldBy[outerBag.name] = outerBag
			outerBag.contains[bagPointer], _ = strconv.Atoi(count)
		}

	}

	return bags
}
