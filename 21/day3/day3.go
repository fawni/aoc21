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
	input = input[:len(input)-1]

	// part one
	var gamma, epilson string

	for i := 0; i < 12; i++ {
		var ones, zeroes int

		for _, binary := range input {
			chars := strings.Split(binary, "")
			num, _ := strconv.Atoi(chars[i])
			switch num {
			case 0:
				zeroes++
			case 1:
				ones++
			}
		}

		if zeroes > ones {
			gamma += "0"
			epilson += "1"
		} else {
			gamma += "1"
			epilson += "0"
		}
	}

	gammadec, _ := strconv.ParseInt(gamma, 2, 64)
	epilsondec, _ := strconv.ParseInt(epilson, 2, 64)
	fmt.Println(gammadec * epilsondec)

	// part two
	common, uncommon := input, input

	for i := 0; i < 12; i++ {
		var ones, zeroes []string
		for _, binary := range common {
			chars := strings.Split(binary, "")
			num, _ := strconv.Atoi(chars[i])
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

	for i := 0; i < 12; i++ {
		var ones, zeroes []string
		for _, binary := range uncommon {
			chars := strings.Split(binary, "")
			num, _ := strconv.Atoi(chars[i])
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
