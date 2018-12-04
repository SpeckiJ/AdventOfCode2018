package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	var twocount int
	var threecount int
	var mapRepo []string

	file, err := os.Open("./day_2/input.file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		// Part 1
		two, three := Two_Three_Count(current)
		if two {
			twocount++
		}
		if three {
			threecount++
		}

		// Part
		mapRepo = append(mapRepo, current)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Result for Part I
	log.Print(twocount * threecount)

	for _, val1 := range mapRepo {
		for _, val2 := range mapRepo {
			if val1 != val2 {
				distance := hamming(val1, val2)
				if distance == 1 {
					// Result for Part II
					log.Print(val2)
					log.Fatal(val1)

				}
			}
		}
	}

}

func Two_Three_Count(input string) (threecount bool, twocount bool) {
	var m = make(map[uint8]int)

	for index := 0; index < len(input); index++ {
		m[input[index]] += 1
	}

	for _, val := range m {
		if val == 2 {
			twocount = true
		}
		if val == 3 {
			threecount = true
		}
	}
	return twocount, threecount
}

func hamming(arg1 string, arg2 string) (int) {

	var distance int
	for i := range arg1 {
		if (arg1[i] != arg2[i]) {
			distance++
		}
	}

	return distance
}