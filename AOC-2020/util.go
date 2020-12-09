package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadData(fileName string) []string {

	file, err := os.Open("input/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	return data
}

// arrayAtoi converts an array of strings into an array of int.
func arrayAtoi(data []string) []int {
	var convertedData = make([]int, len(data))

	for i, _ := range data {
		convertedData[i], _ = strconv.Atoi(data[i])
	}

	return convertedData
}

func splitAtFirst(s string, sep string) (string, string) {
	parts := strings.SplitN(s, sep, 2)

	return parts[0], parts[1]
}
