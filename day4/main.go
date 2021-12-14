package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const board_size int = 5

//go:embed input.txt
var input string

func main() {
	scores := get_scores(input)
	fmt.Printf("First = %d\n", scores[0])
	fmt.Printf("Last = %d\n", scores[len(scores)-1])
}

func get_scores(input string) []int {
	ss := strings.Split(input, "\n")

	numbers := to_ints(strings.Split(ss[0], ","))
	boards := build_boards(ss[2:])
	remaining_boards := [][5][5]int{}

	scores := []int{}
	for _, n := range numbers {
		for _, board := range boards {
			board_complete := false
			for row := range board {
				for col := range board[row] {
					if board[row][col] == n {
						// Board hit - check for full row/column
						board[row][col] = 0
						board_complete = check_row(board[row]) || check_column(col, board)

						if board_complete {
							scores = append(scores, calculate_score(board, n))
						}
					}
				}
			}
			if !board_complete {
				remaining_boards = append(remaining_boards, board)
			}
		}

		boards = remaining_boards
		remaining_boards = nil
	}

	return scores
}

func calculate_score(board [5][5]int, number int) int {
	result := 0
	for _, row := range board {
		for _, n := range row {
			if n != 0 {
				result += n
			}
		}
	}

	return result * number
}

func check_row(row [5]int) bool {
	for _, i := range row {
		if i != 0 {
			return false
		}
	}
	return true
}

func check_column(col int, board [5][5]int) bool {
	for i := 0; i < 5; i++ {
		if board[i][col] != 0 {
			return false
		}
	}
	return true
}

func build_boards(input []string) (boards [][5][5]int) {
	i := 0
	for i < len(input) {
		s := strings.TrimRight(input[i], "\n")
		if s != "" {
			l := input[i : i+board_size]

			board := [5][5]int{}
			for x, r := range l {
				for y, c := range strings.Fields(r) {
					board[x][y] = to_int(c)
				}
			}
			boards = append(boards, board)

			i += board_size
		} else {
			i++
		}
	}
	return boards
}

func to_ints(s []string) (out []int) {
	for _, l := range s {
		out = append(out, to_int(l))
	}
	return out
}

func to_int(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
