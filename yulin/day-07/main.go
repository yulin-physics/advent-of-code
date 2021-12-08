package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	PART_ONE = true
	PART_TWO = false
)

var sum int
var crabPos = []int{}

func main() {
	readInput("input.txt")
	fmt.Println(optimumAlign(crabPos, PART_TWO))
}

func optimumAlign(crabPos []int, constantRate bool) int {
	var pos int
	var fuel int
	if constantRate {
		sort.Ints(crabPos)
		pos = crabPos[len(crabPos)/2]
		for _, v := range crabPos {
			diff := int(math.Abs(float64(v) - float64(pos)))
			fuel += diff
		}
		return fuel
	}
	pos = sum / len(crabPos)
	for _, v := range crabPos {
		diff := stepCost(int(math.Abs(float64(v) - float64(pos))))
		fuel += diff
	}
	return fuel
}

func stepCost(n int) int {
	if n > 0 {
		return n + stepCost(n-1)
	}
	return 0
}

func readInput(fname string) {
	file, err := os.Open(fname)
	if err != nil {

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		for _, r := range row {
			num, err := strconv.Atoi(r)
			if err != nil {
				log.Fatal(err)
			}
			sum += num
			crabPos = append(crabPos, num)
		}
	}
}
