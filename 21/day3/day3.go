package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/x6r/aoc/utils"
)

func main() {
	input := utils.GetInput("input.txt")

	// part one
	var gamma, epsilon string
	bitLen := len(strings.Split(input[0], ""))

	for i := 0; i < bitLen; i++ {
		var ones, zeroes int

		for _, binary := range input {
			num := utils.Atoi(strings.Split(binary, "")[i])
			switch num {
			case 0:
				zeroes++
			case 1:
				ones++
			}
		}

		if zeroes > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammad, _ := strconv.ParseInt(gamma, 2, 64)
	epsilond, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gammad * epsilond)

	// part two
	common, uncommon := input, input

	for i := 0; i < bitLen; i++ {
		var ones, zeroes []string
		for _, binary := range common {
			num := utils.Atoi(strings.Split(binary, "")[i])

			switch num {
			case 0:
				zeroes = append(zeroes, binary)
			case 1:
				ones = append(ones, binary)
			}
		}

		if len(zeroes) > len(ones) {
			common = zeroes
		} else {
			common = ones
		}

		if len(common) == 1 {
			break
		}
	}

	for i := 0; i < bitLen; i++ {
		var ones, zeroes []string
		for _, binary := range uncommon {
			num := utils.Atoi(strings.Split(binary, "")[i])

			switch num {
			case 0:
				zeroes = append(zeroes, binary)
			case 1:
				ones = append(ones, binary)
			}
		}

		if len(zeroes) > len(ones) {
			uncommon = ones
		} else {
			uncommon = zeroes
		}

		if len(uncommon) == 1 {
			break
		}
	}

	oxygen, _ := strconv.ParseInt(common[0], 2, 64)
	scrubber, _ := strconv.ParseInt(uncommon[0], 2, 64)
	fmt.Println(oxygen * scrubber)
}
