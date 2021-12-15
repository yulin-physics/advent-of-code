package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mat struct {
	elements [][]int
}

func main() {
	m := readInput("input.txt")
	fmt.Println(m)
}

func readInput(fname string) Mat {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := Mat{}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		nums := []int{}
		for _, r := range row {
			num, err := strconv.Atoi(r)
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		m.elements = append(m.elements, nums)
	}
	return m
}
