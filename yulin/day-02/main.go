package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

func main() {
	data, err := readStr("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part one: %d\npart two: %d", crossSec(data), crossSecWithAim(data))

}

func crossSec(data string) int {
	reDir := regexp.MustCompile(`(forward|down|up)`)
	directions := reDir.FindAllString(data, -1)
	reDis := regexp.MustCompile(`\d+`)
	distances := reDis.FindAllString(data, -1)
	pos := [2]int{0, 0}
	for i, direction := range directions {
		n, _ := strconv.Atoi(distances[i])
		switch direction {
		case "forward":
			pos[1] += n
		case "down":
			pos[0] += n
		case "up":
			pos[0] -= n
		}
	}
	return pos[0] * pos[1]
}

func crossSecWithAim(data string) int {
	reDir := regexp.MustCompile(`(forward|down|up)`)
	directions := reDir.FindAllString(data, -1)
	reDis := regexp.MustCompile(`\d+`)
	distances := reDis.FindAllString(data, -1)
	pos := [2]int{0, 0}
	var aim int
	for i, direction := range directions {
		n, _ := strconv.Atoi(distances[i])
		switch direction {
		case "forward":
			pos[1] += n
			pos[0] += aim * n
		case "down":
			aim += n
		case "up":
			aim -= n
		}
	}
	return pos[0] * pos[1]
}

func readStr(fname string) (string, error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
