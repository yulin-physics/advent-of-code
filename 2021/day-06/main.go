package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type School struct {
	fish      []int
	countFish []int
}

var states = []int{6, 0, 1, 2, 3, 4, 5, 6, 7, 8}

func main() {
	school := readInput("input.txt")
	fmt.Println(school.calcSchool(256))
}

//works for large days, calc number of fish for each fish age/internal timer
func (s *School) calcSchool(days int) int {
	for range make([]struct{}, days) {
		s.nextDay()
	}
	var total int
	for _, c := range s.countFish {
		total += c
	}
	return total
}

func (s *School) nextDay() {
	newFish := s.countFish[0]
	s.countFish = append(s.countFish[1:], newFish)
	s.countFish[6] += newFish
}

//works for 80 days, simulates fish age
func (s *School) simulateFish(days int) int {
	for range make([]struct{}, days) {
		for i, f := range s.fish {
			s.fish[i] = states[f]
			if f == 0 {
				s.fish = append(s.fish, states[9])
			}
		}
	}
	return len(s.fish)
}

func readInput(fname string) School {
	file, err := os.Open(fname)
	if err != nil {

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	countFish := make([]int, 9)
	fish := []int{}
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		for _, r := range row {
			num, err := strconv.Atoi(r)
			if err != nil {
				log.Fatal(err)
			}
			countFish[num]++
			fish = append(fish, num)
		}
	}
	return School{fish: fish, countFish: countFish}
}
