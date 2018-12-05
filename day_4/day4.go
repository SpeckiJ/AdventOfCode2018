package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Observation struct {
	Month int
	Day int
	Hours int
	Minutes int
	Obs string
}

type Timer struct {
	TimeAsleep int
	Times map[int]int
}

func main() {
	file, err := os.Open("./day_4/input.file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []Observation

	// PART I
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		var obs Observation
		// [1518-04-21 00:02] Guard #3361 begins shift
		split := strings.Split(current, "]")
		datetime := strings.Split(split[0], " ")
		date_seq := strings.Split(datetime[0], "-")
		time_seq := strings.Split(datetime[1], ":")

		obs.Month, _ = strconv.Atoi(date_seq[1])
		obs.Day, _ = strconv.Atoi(date_seq[2])
		obs.Hours, _ = strconv.Atoi(time_seq[0])
		obs.Minutes, _ = strconv.Atoi(time_seq[1])
		obs.Obs = split[1][1:]

		lines = append(lines, obs)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort in-place
	quicksort(lines, 0, len(lines)-1)

	// PART I
	m := asleepTimer(lines)
	max_key := 0
	max_val := 0
	for key, value := range m {
		if value.TimeAsleep > max_val {
			max_key = key
			max_val = value.TimeAsleep
		}
	}

	max_timestamp := 0
	max_timestamp_key := 0
	for l, value := range m[max_key].Times {
		if value > max_timestamp {
			max_timestamp_key = l
			max_timestamp = value
		}
	}
	log.Print(max_key * max_timestamp_key)

	// PART II

	max_timestamp_guard_key := 0
	max_timestamp_count := 0
	max_timestamp_key = 0
	for key := range m {
		for l, value := range m[key].Times {
			if value >= max_timestamp_count {
				max_timestamp_key = l
				max_timestamp_count = value
				max_timestamp_guard_key = key
			}
		}
	}
	fmt.Println(max_timestamp_count)
	fmt.Println(max_timestamp_key)
	fmt.Println(max_timestamp_guard_key)
}

func asleepTimer(input []Observation) map[int]*Timer {
	guards := make(map[int]*Timer)
	var current_guard int
	var asleepTime int
	var asleep bool

	for _, val := range input {
		split := strings.Split(val.Obs, " ")
		// new Guard
		if (len(split)) > 2 {
			asleepTime = 0
			asleep = false
			current_guard, _= strconv.Atoi(split[1][1:])
			if _, ok := guards[current_guard]; !ok {
				guards[current_guard] = &Timer{0, make(map[int]int)}
			}
			continue
		}
		// Old Guard
		if asleep {
			// Record asleep duration
			minutesAsleep := val.Minutes - asleepTime
			guards[current_guard].TimeAsleep += minutesAsleep

			// Record which Times guard is asleep
			for i := asleepTime; i < val.Minutes; i++ {
				guards[current_guard].Times[i] += 1
			}
			asleep = false
		} else {
			asleep = true
			asleepTime = val.Minutes
		}
	}
	return guards
}

func quicksort(input []Observation, low, high int) {
	if low < high {
		pivot := quicksort_rec(input, low, high)
		quicksort(input, low, pivot -1)
		quicksort(input, pivot+1, high)
	}
}

func quicksort_rec(input []Observation, low, high int) int {
	pivot := high
	i := low -1

	for j := low; j < high; j++ {
		if compare(input[j], input[pivot]) < 0  {
			i++
			tmp :=input[j]
			input[j] = input[i]
			input[i] = tmp
		}
	}
	tmp := input[i+1]
	input[i+1] = input[high]
	input[high] = tmp
	return i+1
}

func compare(arg1, arg2 Observation) int {
	if arg1.Month > arg2.Month {
		return 1
	} else if arg1.Month < arg2.Month {
		return -1
	}
	if arg1.Day > arg2.Day {
		return 1
	} else if arg1.Day < arg2.Day {
		return -1
	}
	if arg1.Hours > arg2.Hours {
		return 1
	} else if arg1.Hours < arg2.Hours {
		return -1
	}
	if arg1.Minutes > arg2.Minutes {
		return 1
	} else if arg1.Minutes < arg2.Minutes {
		return -1
	}
	return 0
}
