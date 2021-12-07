package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const(
	PART_ONE = false
	PART_TWO = true
)
var points = map[[2]int]int{}

func main() {
	lines := readInput("input.txt")
	fmt.Println(numOfOverlapPoints(lines, PART_TWO))
}

func numOfOverlapPoints(lines [][]int, includeDiagonal bool) int {
	for _, line := range lines {
		if line[0] == line[2] {
			for _, y := range makeRange(line[1], line[3]) {
				points[[2]int{line[0], y}] += 1
			}
		} else if line[1] == line[3] {
			for _, x := range makeRange(line[0], line[2]) {
				points[[2]int{x, line[1]}] += 1
			}
		} else if includeDiagonal && math.Abs(float64(line[0]-line[2])) == math.Abs(float64(line[1]-line[3])) {
			x := makeRange(line[0], line[2])
			y := makeRange(line[1], line[3])
			for i := range x {
				points[[2]int{x[i], y[i]}] += 1
			}
		}
	}
	return len(points) - len(reverseMap(points)[1])
}

func makeRange(min, max int) []int {
	if min > max {
		slice := make([]int, min-max+1)
		for i := range slice {
			slice[i] = min - i
		}
		return slice
	}
	slice := make([]int, max-min+1)
	for i := range slice {
		slice[i] = min + i
	}
	return slice
}

func reverseMap(points map[[2]int]int) map[int][][2]int {
	rev := map[int][][2]int{}
	for k, v := range points {
		rev[v] = append(rev[v], k)
	}
	return rev
}

func readInput(fname string) [][]int {
	file, err := os.Open(fname)
	if err != nil {

	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		pair := []int{}
		for _, r := range row {
			p, _ := strconv.Atoi(r)
			pair = append(pair, p)
		}
		lines = append(lines, pair)
	}
	return lines
}
