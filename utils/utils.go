package utils

import (
	"os"
	"strconv"
	"strings"
)

func GetInput(path string) []string {
	file, _ := os.ReadFile(path)
	input := strings.Split(string(file), "\n")
	input = input[:len(input)-1]
	return input
}

func Atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
