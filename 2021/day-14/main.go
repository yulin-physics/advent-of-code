package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Polymerisation struct {
	template string
	rules    map[string]string
	allRules string
	elements []string
}

func main() {
	p := readInput("input.txt")
	p.growPolymer(40)
}

func (p *Polymerisation) growPolymer(steps int) int {
	for range make([]struct{}, steps) {
		p.template = p.pairInsertion(p.template, "")
	}
	max := strings.Count(p.template, p.elements[0])
	min := strings.Count(p.template, p.elements[0])
	for _, el := range p.elements {
		num := strings.Count(p.template, el)
		if num > max {
			max = num
		} else if num < min {
			min = num
		}
	}
	fmt.Println(max - min)
	return max - min
}

func (p *Polymerisation) pairInsertion(s string, new string) string {
	if len(s) == 2 && strings.Contains(p.allRules, s) {
		new += p.rules[s]
	}else if strings.Contains(p.allRules, s[:2]) {
		new += p.rules[s[:2]][:2]
		return p.pairInsertion(s[1:], new)
	}
	return new
}

func readInput(fname string) Polymerisation {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	p := Polymerisation{}
	p.rules = make(map[string]string)
	start := true
	for scanner.Scan() {
		if start {
			p.template = scanner.Text()
			start = false
		} else if scanner.Text() == "" {
			continue
		} else {
			rule := strings.Split(scanner.Text(), " -> ")
			temp := ""
			for i, r := range rule[0] {
				if i == 1 {
					temp += rule[1]
				}
				temp += string(r)
			}
			p.rules[rule[0]] = temp
			p.allRules += fmt.Sprintf("%s, ", rule[0])
		}
	}
	p.elements = []string{"B", "C", "H", "N"}
	return p
}
