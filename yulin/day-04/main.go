package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	boards, randNums, _ := readInput("input.txt")
	findWinners(boards, randNums)
	fmt.Printf("part one: %d, part two: %d", winners["first"].finalScore(), winners["last"].finalScore())
}

func findFirstWinner(boards []*Board, randNums []int) (boardInd int, updateRandNums []int) {
	if len(randNums) == 0 {
		return -1, nil
	}
	for b, board := range boards {
		for i, row := range board.Row {
			for j, num := range row {
				if num == randNums[0] {
					board.Row[i][j] = -1
					board.calledNums = append(board.calledNums, num)
				}
			}
		}
		if hasWon(board.Row) {
			//if first winner only: change key="first", then call this func in main
			winners["last"] = board
			return b, randNums
		}
	}
	return findFirstWinner(boards, randNums[1:])
}

func findWinners(boards []*Board, randNums []int) (boardInt int, updateRandNums []int) {
	first := true
	for range randNums {
		if first {
			boardInt, randNums = findFirstWinner(boards, randNums)
			winners["first"] = boards[boardInt]
			first = false
		}
		boards = append(boards[:boardInt], boards[boardInt+1:]...)
		boardInt, randNums = findFirstWinner(boards, randNums)
		//no more winners found
		if randNums == nil {
			break
		}
	}
	return boardInt, randNums
}

func (b *Board) finalScore() int {
	var unmarked int
	for _, row := range b.Row {
		for _, num := range row {
			if num != -1 {
				unmarked += num
			}
		}
	}
	return unmarked * b.calledNums[len(b.calledNums)-1]
}

func hasWon(rows [5][5]int) bool {
	for i, row := range rows {
		if (row[0] == -1 && row[1] == -1 && row[2] == -1 && row[3] == -1 && row[4] == -1) ||
			(rows[0][i] == -1 && rows[1][i] == -1 && rows[2][i] == -1 && rows[3][i] == -1 && rows[4][i] == -1) {
			return true
		}
	}
	return false
}

func readInput(fname string) (boards []*Board, randNum []int, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	boardRow := [5][5]int{}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		read_line := scanner.Text()
		if i == 0 {
			randNum = stringtoInt(strings.Split(read_line, ","))
			i++
			continue
		} else if len(read_line) == 0 {
			continue
		}
		row := strings.Fields(read_line)
		for j, n := range row {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, nil, err
			}
			boardRow[(i-1)%5][j] = num
		}
		if i%5 == 0 {
			boards = append(boards, &Board{Row: boardRow})
		}
		i++
	}
	return boards, randNum, nil
}

func stringtoInt(slice []string) []int {
	var ints []int
	for _, s := range slice {
		n, _ := strconv.Atoi(s)
		ints = append(ints, n)
	}
	return ints
}

var winners = map[string]*Board{"first": {}, "last": {}}

type Board struct {
	Row        [5][5]int
	calledNums []int
}
