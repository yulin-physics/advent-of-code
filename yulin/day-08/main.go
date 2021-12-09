package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var inputDigits = [][]string{}

func main() {
	readInput("input.txt")
	fmt.Printf("part one: %d\npart two: %d", count1478(inputDigits), decodeRows(inputDigits))
}

func decodeRows(digits [][]string) int {
	var total int
	for _, row := range digits {
		segToDigit := decodeDigits(row[:10])
		value := segToDigit[row[10]] + segToDigit[row[11]] + segToDigit[row[12]] + segToDigit[row[13]]
		num, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		total += num
	}
	return total
}

func decodeDigits(segments []string) map[string]string {
	digitToSeg := make([]string, 10)
	segToDigit := make(map[string]string, 10)
	var fiveAndSixSeg = map[int][]string{}
	for _, seg := range segments {
		switch len(seg) {
		case 2:
			digitToSeg[1], segToDigit[seg] = seg, "1"
		case 4:
			digitToSeg[4], segToDigit[seg] = seg, "4"
		case 3:
			digitToSeg[7], segToDigit[seg] = seg, "7"
		case 7:
			digitToSeg[8], segToDigit[seg] = seg, "8"
		//len of 5: digits 2, 3, 5
		//len of 6: digits 0, 6, 9
		default:
			fiveAndSixSeg[len(seg)] = append(fiveAndSixSeg[len(seg)], seg)
		}
	}
	segToDigit["topBLbottom"] = remove(digitToSeg[8], digitToSeg[4])
	digitToSeg, segToDigit = decode235(fiveAndSixSeg[5], digitToSeg, segToDigit)
	digitToSeg, segToDigit = decode069(fiveAndSixSeg[6], digitToSeg, segToDigit)
	return segToDigit
}

func decode235(segments []string, digitToSeg []string, segToDigit map[string]string) ([]string, map[string]string) {
	for _, seg := range segments {
		if contains(seg, digitToSeg[1]) {
			digitToSeg[3], segToDigit[seg] = seg, "3"
		} else if contains(seg, segToDigit["topBLbottom"]) {
			digitToSeg[2], segToDigit[seg] = seg, "2"
		} else {
			digitToSeg[5], segToDigit[seg] = seg, "5"
		}
	}
	return digitToSeg, segToDigit
}

func decode069(segments []string, digitToSeg []string, segToDigit map[string]string) ([]string, map[string]string) {
	for _, seg := range segments {
		if !contains(seg, digitToSeg[1]) {
			digitToSeg[6], segToDigit[seg] = seg, "6"
		} else if contains(seg, digitToSeg[4]) {
			digitToSeg[9], segToDigit[seg] = seg, "9"
		} else {
			digitToSeg[0], segToDigit[seg] = seg, "0"
		}
	}
	return digitToSeg, segToDigit
}

func contains(s1, s2 string) bool {
	var count int
	for _, x := range s1 {
		for _, y := range s2 {
			if x == y {
				count++
			}
		}
	}
	return count == len(s2)
}

func remove(s1, s2 string) string {
	var s string
next:
	for _, x := range s1 {
		for _, y := range s2 {
			if x == y {
				continue next
			}
		}
		s += string(x)
	}
	return s
}

func count1478(digits [][]string) int {
	var count int
	for _, row := range digits {
		for _, digit := range row[10:] {
			if len(digit) == 2 || len(digit) == 4 || len(digit) == 3 || len(digit) == 7 {
				count++
			}
		}
	}
	return count
}

func readInput(fname string) {
	file, err := os.Open(fname)
	if err != nil {

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.FieldsFunc(scanner.Text(), func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		var sorted = []string{}
		for _, s := range row {
			w := strings.Split(s, "")
			sort.Strings(w)
			sorted = append(sorted, strings.Join(w, ""))
		}
		inputDigits = append(inputDigits, sorted)
	}
}
