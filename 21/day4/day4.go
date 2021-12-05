package main

import (
	"fmt"
	"strings"

	"github.com/x6r/aoc/utils"
)

type board struct {
	numbers [5][5]int
	marked  [5][5]bool
}

func main() {
	input := utils.GetInput("input.txt")
	drawnNumbers := strings.Split(input[0], ",")
	boards := parseBoards(input)
	winningBoard, winningNumber := 0, 0

	// part one
outerone:
	for _, drawn := range drawnNumbers {
		mark(utils.Atoi(drawn), boards)
		for idx, board := range boards {
			if bingo(board) {
				winningBoard, winningNumber = idx, utils.Atoi(drawn)
				break outerone
			}
		}
	}
	fmt.Println(score(boards[winningBoard], winningNumber))

	// part two
	freshBoards := parseBoards(input)
	var lastWon int
	won := make(map[int]bool)
outertwo:
	for _, drawn := range drawnNumbers {
		mark(utils.Atoi(drawn), freshBoards)
		for idx, board := range freshBoards {
			if won[idx] {
				continue
			}
			if bingo(board) {
				won[idx] = true
				lastWon, winningNumber = idx, utils.Atoi(drawn)
				if len(won) == len(boards) {
					break outertwo
				}
			}
		}
	}
	fmt.Println(score(freshBoards[lastWon], winningNumber))
}

func parseBoards(input []string) []board {
	boardsCounter, rowIndex := 0, 0
	boards := make([]board, len(input)/5-(len(input)/5-len(input)/6))

	for _, row := range input[2:] {
		if row == "" {
			rowIndex = 0
			boardsCounter++
			continue
		}
		for numberIndex, number := range strings.Fields(row) {
			num := utils.Atoi(number)
			boards[boardsCounter].numbers[rowIndex][numberIndex] = num
		}
		rowIndex++
	}
	return boards
}

func mark(drawn int, boards []board) {
	for idx, board := range boards {
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				if board.numbers[row][col] == drawn {
					boards[idx].marked[row][col] = true
				}
			}
		}
	}
}

func bingo(board board) bool {
rows:
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !board.marked[row][col] {
				continue rows
			}
		}
		return true
	}
columns:
	for col := 0; col < 5; col++ {
		for row := 0; row < 5; row++ {
			if !board.marked[row][col] {
				continue columns
			}
		}
		return true
	}
	return false
}

func score(board board, winningNumber int) int {
	sum := 0
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if !board.marked[row][col] {
				sum += board.numbers[row][col]
			}
		}
	}
	return sum * winningNumber
}
