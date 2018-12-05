package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./day_5/input.file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input string
	// PART I
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		if current != "" {
			input = current
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	characters := []rune(input)

	minimum := len(characters)
	var minimum_rune rune
	for _, val := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		len := reactPolymer([]rune(input), val)
		if len < minimum {
			minimum = len
			minimum_rune = val
		}
	}

	log.Print(len(input))
	log.Print(minimum)
	log.Print(minimum_rune)
}

func reactPolymer(characters []rune, ignoredRunes rune) int {
	changed := true
	size := len(characters)
	// remove ignored runes
	for j := range characters {
		if characters[j] == ignoredRunes || characters[j] == (ignoredRunes + 32)  {
			size--
			characters[j] = 0
		}
	}

	// Iterate until nothing changes
	for changed {
		changed = false
		last_index := 0
		for i := range characters {
			if characters[i] != 0 {
				if characters[i] - characters[last_index] != 32 && characters[i] - characters[last_index] != -32 {
					// Advance last and keep everything
					last_index = i
				} else {
					// Delete both characters and continue
					characters[last_index] = 0
					characters[i] = 0
					size -= 2
					changed = true
				}
			}
		}
	}
	return size
}
