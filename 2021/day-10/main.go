package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var pairs = map[string]string{
	"(": ")",
	"<": ">",
	"{": "}",
	"[": "]",
}

var closingErrorScore = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var closingCompletionScore = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

var brackets = [][]string{}

func main() {
	readInput("input.txt")
	errorScore, brackets := syntaxErrorScore(brackets)
	completionScore := completionMiddleScore(brackets)
	fmt.Printf("part one: %d\npart two: %d", errorScore, completionScore)
}

func completionMiddleScore(brackets [][]string) int {
	var scores []int
	for _, row := range brackets {
		var score int
		closings := removePairs(row)
		for i := len(closings) - 1; i >= 0; i-- {
			score = score*5 + closingCompletionScore[pairs[closings[i]]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func syntaxErrorScore(brackets [][]string) (int, [][]string) {
	var score int
	var cleanBrackets [][]string
next:
	for _, row := range brackets {
		for _, bracket := range removePairs(append([]string(nil), row...)) {
			if closingErrorScore[bracket] != 0 {
				score += closingErrorScore[bracket]
				continue next
			}
		}
		cleanBrackets = append(cleanBrackets, row)
	}
	return score, cleanBrackets
}

func removePairs(brackets []string) []string {
	for i := 1; i < len(brackets); i++ {
		if pairs[brackets[i-1]] == brackets[i] {
			brackets := append(brackets[:i-1], brackets[i+1:]...)
			return removePairs(brackets)
		}
	}
	return brackets
}

func readInput(fname string) {
	file, err := os.Open(fname)
	if err != nil {

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		brackets = append(brackets, row)
	}
}
