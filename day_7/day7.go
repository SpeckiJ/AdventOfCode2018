package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Step struct {
	Name string
	Requires []string
	Time int
}

type Worker struct {
	Status bool
	Task Step
}

	func main() {
		file, err := os.Open("./day_7/input.file")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var input []string
		// PART I
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			current := scanner.Text()
			if err != nil {
				log.Fatal(err)
			}
			if current != "" {
				input = append(input, current)
			}
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		graph := parseToGraph(input)


		var output string
		doneList := make(map[string]bool)
		for len(graph) > 0 {
			// get possible steps
			var possible []*Step
			for _, v := range graph {
				allDone := true
				for _, val := range v.Requires {
					if _, ok := doneList[val]; !ok {
						allDone = false
					}
				}
				if allDone {
					possible = append(possible, v)
				}
			}

			// Check alphabetically for execution order
			first := possible[0]
			for _, v := range possible {
				if first.Name > v.Name {
					first = v
				}
			}
			output += first.Name
			delete(graph, first.Name)
			doneList[first.Name] = true
		}

		log.Print(output)

		// Part II
		doneList = make(map[string]bool)
		workers := [5]*Worker{
			{false, Step{"", make([]string, 0), 0}},
			{false, Step{"", make([]string, 0), 0}},
			{false, Step{"", make([]string, 0), 0}},
			{false, Step{"", make([]string, 0), 0}},
			{false, Step{"", make([]string, 0), 0}},
		}
		output = ""

		graph = parseToGraph(input)
		elements := len(graph)
		time := -1
		for elements > 0 {
			time++
			for _, worker := range workers {
				// Clear up old Jobs if completed
				if worker.Status {
					if worker.Task.Time == 0 {
						// Available for next Job
						worker.Status = false
						output += worker.Task.Name
						elements--
						doneList[worker.Task.Name] = true
					}
					worker.Task.Time--
				}
				if !worker.Status {
					// Assign new Job
					// get possible steps
					var possible []*Step
					for _, v := range graph {
						allDone := true
						for _, val := range v.Requires {
							if _, ok := doneList[val]; !ok {
								allDone = false
							}
						}
						if allDone {
							possible = append(possible, v)
						}
					}

					// If Step is available assign to worker
					if possible != nil {
						// Check alphabetically for execution order
						first := possible[0]
						for _, v := range possible {
							if first.Name > v.Name {
								first = v
							}
						}
						worker.Status = true
						worker.Task = *first
						delete(graph, worker.Task.Name)
						worker.Task.Time--
					}
				}
			}
			fmt.Print(time)
		}

		log.Print(output)
		log.Fatal(time)
	}


func parseToGraph(input []string) map[string]*Step {
	// Parse to struct
	graph := make(map[string]*Step)

	for _, val := range input {
		split := strings.Split(val, " ")

		// Init with Name if not initialized
		if _, ok := graph[split[1]]; !ok {
			graph[split[1]] = &Step{split[1], make([]string, 0), int(split[1][0])-4}
		}
		if _, ok := graph[split[7]]; !ok {
			graph[split[7]] = &Step{split[7], make([]string, 0), int(split[7][0])-4}
		}
		graph[split[7]].Requires = append(graph[split[7]].Requires, split[1])
	}
	return graph
}
