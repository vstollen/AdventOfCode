package main

import (
	"bufio"
	"log"
	"os"
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
