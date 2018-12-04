package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day_1/input.file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var operations []int
	var found = false
	var previousResults []int
	var previous = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		operations = append(operations, current)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for !found {
		for op := range operations {
			previous += operations[op]
			log.Print(previous)
			log.Print(op)
			if contains(previousResults, previous) {
				found = true
				break
			}
			previousResults = append(previousResults, previous)
		}
	}

	log.Fatal(previous)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
