package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	metadata []int
	children []Node
}

func main() {
	file, err := os.Open("./day_8/input.file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var raw string

	// PART I
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		if current != "" {
			raw = current
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	split := strings.Split(raw, " ")
	input := make([]int, len(split))
	for k, v := range split {
		input[k], _ = strconv.Atoi(v)
	}
	root, _ := parseNode(input)


	log.Print(addMetadataSimple(root))

	log.Print(addMetadataComplex(root))
}

func addMetadataComplex(root Node) (count int) {
	// Simply add Metadata if no childs are present
	if len(root.children) == 0 {
		for _, k := range root.metadata {
			count += k
		}
		return count
	} else {
		for _ , v := range root.metadata {
			if v == 0 || len(root.children) < v{
				continue
			}
			count += addMetadataComplex(root.children[v-1])
		}
	}
	return count
}

func addMetadataSimple(root Node) (count int) {
	for _, k := range root.metadata {
		count += k
	}
	for _, v := range root.children {
		count += addMetadataSimple(v)
	}
	return count
}

func parseNode(input []int) (result Node, index int) {
	childNotes := input[0]
	metadata := input[1]
	result.children = make([]Node, childNotes)
	result.metadata = make([]int, metadata)
	// Ingested two numbers already
	index += 2

	var ingested int
	for i := 0; i < childNotes; i++ {
		result.children[i], ingested = parseNode(input[index:])
		index += ingested
	}

	for i := 0; i < metadata; i++ {
		result.metadata[i] = input[index]
		index++
	}
	return result, index
}
