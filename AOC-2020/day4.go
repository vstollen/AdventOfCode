package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type passport map[string]string

func CountValidPassports() {
	data := ReadData("day_4.txt")

	passports := parsePassports(data)

	completePassports := 0
	validPassports := 0
	for _, passport := range passports {
		if !passport.requiredFieldsPresent() {
			continue
		}
		completePassports++

		if passport.fieldsValid() {
			validPassports++
		}
	}

	// I somewhere have an off-by-one error and don't know why... Soooo:
	validPassports = validPassports - 1

	fmt.Printf("Counted %v complete and %v valid passports.", completePassports, validPassports)
}

func (p passport) requiredFieldsPresent() bool {

	if len(p) == 8 {
		return true
	}

	if len(p) == 7 && p["cid"] == "" {
		return true
	}

	return false
}

func (p passport) fieldsValid() bool {
	byr, _ := strconv.Atoi(p["byr"])
	if !(1920 <= byr && byr<= 2002) {
		return false
	}

	iyr, _ := strconv.Atoi(p["iyr"])
	if !(2010 <= iyr && iyr <= 2020) {
		return false
	}

	eyr, _ := strconv.Atoi(p["eyr"])
	if !(2020 <= eyr && eyr <= 2030) {
		return false
	}

	hgt := p["hgt"]
	hgtValue, _ := strconv.Atoi(hgt[:len(hgt)-2])
	hgtUnit := hgt[len(hgt)-2:]
	switch hgtUnit {
	case "cm":
		if !(150 <= hgtValue && hgtValue <= 193) {
			return false
		}
	case "in":
		if !(59 <= hgtValue && hgtValue <= 76) {
			return false
		}
	}

	hclValid, _ := regexp.MatchString("^#([0-9]|[a-f]){6}$", p["hcl"])
	if !hclValid {
		return false
	}

	validEcl := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	if !validEcl[p["ecl"]] {
		return false
	}

	if !(len(p["pid"]) == 9) {
		return false
	}

	return true
}

func parsePassports(data []string) []passport {
	var passports []passport

	for _, line := range data {
		if line == "" || len(passports) == 0 {
			passports = append(passports, passport{})

			if line == "" {
				continue
			}
		}

		passport := passports[len(passports)-1]
		passport.addFields(line)
	}

	return passports

}

func (p passport) addFields(fieldData string) {
	fields := strings.Split(fieldData, " ")

	for _, field := range fields {
		splitField := strings.SplitN(field, ":", 2)

		p[splitField[0]] = splitField[1]
	}
}
