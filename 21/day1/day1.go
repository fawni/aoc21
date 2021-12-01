package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.Split(string(file), "\n")
	input = input[:len(input)-1] // there is an additional space in the last line of the input file which messes with part two's sum

	var increased int
	current, _ := strconv.Atoi(input[0])

	// part one
	for _, s := range input[1:] {
		new, _ := strconv.Atoi(s)
		if new > current {
			increased++
		}
		current = new
	}

	fmt.Println(increased)

	// part two
	increased = 0
	sum := 0
	nums := []int{0, 0, 0}
	for idx, s := range input {
		num, _ := strconv.Atoi(s)
		currentSum := sum
		sum += num - nums[0]
		nums = append(nums[1:], num)
		if idx < 3 {
			continue
		}
		if sum > currentSum {
			increased++
		}
	}

	fmt.Println(increased)
}
