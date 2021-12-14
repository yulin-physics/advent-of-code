package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Polymerisation struct {
	template string
	polymer  map[string]int
	rules    map[string]string
}

func main() {
	p := readInput("input.txt")
	fmt.Println(p.growPolymer(40))
}

func (p *Polymerisation) growPolymer(steps int) int {
	for range make([]struct{}, steps) {
		p.pairInsertion(p.polymer)
	}
	counts := p.countElements()
	return counts[len(counts)-1] - counts[0]
}

func (p *Polymerisation) pairInsertion(template map[string]int) {
	new := map[string]int{}
	for k, v := range template {
		second := p.rules[k]
		new[string(k[0])+second] += v
		new[second+string(k[1])] += v
	}
	p.polymer = new
}

func (p *Polymerisation) countElements() []int {
	elements := map[string]int{}
	//double counting all elements
	elements[string(p.template[0])]++
	elements[string(p.template[len(p.template)-1])]++
	for k, v := range p.polymer {
		elements[string(k[0])] += v
		elements[string(k[1])] += v
	}
	counts := []int{}
	for _, v := range elements {
		counts = append(counts, v/2)
	}
	sort.Ints(counts)
	return counts
}

func readInput(fname string) Polymerisation {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	p := Polymerisation{}
	p.rules, p.polymer = make(map[string]string), make(map[string]int)
	start := true
	for scanner.Scan() {
		if start {
			p.template = scanner.Text()
			for i := 1; i < len(p.template); i++ {
				p.polymer[p.template[i-1:i+1]]++
			}
			start = false
		} else if scanner.Text() == "" {
			continue
		} else {
			rule := strings.Split(scanner.Text(), " -> ")
			p.rules[rule[0]] = rule[1]
		}
	}
	return p
}
