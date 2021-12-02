package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(file), "\n")
	input = input[:len(input)-1]

	var hpos, depth int

	// part one
	for _, line := range input {
		slice := strings.Fields(line)
		instruction := slice[0]
		amount, _ := strconv.Atoi(slice[1])
		switch instruction {
		case "forward":
			hpos += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}

	fmt.Println(hpos * depth)

	// part two
	hpos, depth = 0, 0
	var aim int

	for _, line := range input {
		slice := strings.Fields(line)
		instruction := slice[0]
		amount, _ := strconv.Atoi(slice[1])
		switch instruction {
		case "down":
			aim += amount
		case "up":
			aim -= amount
		case "forward":
			hpos += amount
			depth += aim * amount
		}
	}

	fmt.Println(hpos * depth)
}
