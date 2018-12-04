package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	Id string
	Start_x int
	Start_y int
	Size_x int
	Size_y int
}

func main() {
	var canvas [1000][1000]int
	var operations []Operation
	file, err := os.Open("./day_3/input.file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// PART I
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		split := strings.Split(current, " ")
		start := strings.Split(split[2][:len(split[2])-1], ",")
		size := strings.Split(split[3], "x")

		start_x, _ := strconv.Atoi(start[0])
		start_y, _ := strconv.Atoi(start[1])

		size_x, _ := strconv.Atoi(size[0])
		size_y, _ := strconv.Atoi(size[1])
		operations = append(operations, Operation{split[0], start_x, start_y, size_x, size_y})

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Setup canvas
	for _, current := range operations {
		for i := 0; i < current.Size_x; i++ {
			for j := 0; j < current.Size_y; j++ {
				canvas[current.Start_x+i][current.Start_y+j]++
			}
		}
	}
	// Compute Overlapping Number
	var overlapping int
	for i := 0; i < len(canvas); i++ {
		for j := 0; j < len(canvas[0]); j++ {
			if canvas[i][j] >= 2 {
				overlapping++
			}
		}
	}
	log.Print(overlapping)

	// PART II
	for _, current := range operations {
		var claimed bool
		// Check if any operation has only single claimed squares
		for i := 0; i < current.Size_x; i++ {
			for j := 0; j < current.Size_y; j++ {
				if canvas[current.Start_x+i][current.Start_y+j] > 1 {
					// Claimed by somebody else
					claimed = true
				}
			}
		}
		if !claimed {
			log.Print(current.Id)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
