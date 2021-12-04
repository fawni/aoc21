package main

import (
	"fmt"

	"github.com/x6r/aoc/utils"
)

func main() {
	input := utils.GetInput("input.txt")

	var increased int
	current := utils.Atoi(input[0])

	// part one
	for _, new := range input[1:] {
		if utils.Atoi(new) > current {
			increased++
		}
		current = utils.Atoi(new)
	}

	fmt.Println(increased)

	// part two
	increased = 0
	nums, sum := []int{0, 0, 0}, 0

	for idx, s := range input {
		num := utils.Atoi(s)
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
