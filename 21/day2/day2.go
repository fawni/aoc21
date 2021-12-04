package main

import (
	"fmt"
	"strings"

	"github.com/x6r/aoc/utils"
)

func main() {
	input := utils.GetInput("input.txt")

	var hpos, depth int

	// part one
	for _, line := range input {
		slice := strings.Fields(line)
		instruction, amount := slice[0], utils.Atoi(slice[1])
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
		instruction, amount := slice[0], utils.Atoi(slice[1])
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
