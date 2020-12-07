package main

import (
	"fmt"
	"strings"
)

type bag struct {
	name string
	heldBy map[string]*bag
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

	fmt.Printf("Total: %v", len(outerBags))
}

func appendOuterParents(innerBag *bag, outerBags map[*bag]bool) map[*bag]bool {
	outerBags[innerBag] = true

	for _, outerBag := range innerBag.heldBy {
		outerBags = appendOuterParents(outerBag, outerBags)
	}

	return outerBags
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
		}
		outerBag = bags[bagName]

		rule = strings.TrimSuffix(rule, ".")

		rules := strings.Split(rule, ", ")

		for _, rule := range rules {
			rule = strings.TrimSuffix(rule, " bag")
			rule = strings.TrimSuffix(rule, " bags")

			_, name := splitAtFirst(rule, " ")

			bagPointer, ok := bags[name]
			if !ok {
				bags[name] = &bag{
					name:   name,
					heldBy: map[string]*bag{},
				}
				bagPointer = bags[name]
				bagPointer.heldBy = make(map[string]*bag)
			}

			bagPointer.heldBy[outerBag.name] = outerBag
		}

	}

	return bags
}
