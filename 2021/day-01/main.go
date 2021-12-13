package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	nums, err := readInts("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//part one
	pOne := numOfIncreasedMeasurement(nums, 1)
	//part two
	pTwo := numOfIncreasedMeasurement(nums, 3)
	fmt.Printf("part one: %d\npart two: %d", pOne, pTwo)
}

func numOfIncreasedMeasurement(nums []int, pos int) int {
	var count int
	for i := pos; i < len(nums); i++ {
		if nums[i] > nums[i-pos] {
			count++
		}
	}
	return count
}

func readInts(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(b), "\n")
	nums = make([]int, 0, len(lines))
	for _, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}
