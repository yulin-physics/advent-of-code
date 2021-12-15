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

	"../utils"
)

type Element struct {
	Position [2]int
	Priority int
}

type PriorityQueue struct {
	Elements []Element
}
type Mat struct {
	Elements [][]int
}

func (pq *PriorityQueue) sort() {
	sort.SliceStable(pq.Elements, func(i, j int) bool {
		return pq.Elements[i].Priority > pq.Elements[j].Priority
	})
}

func (pq *PriorityQueue) pop() Element {
	pq.sort()
	top := pq.Elements[0]
	pq.Elements = pq.Elements[1:]
	return top
}

func (pq *PriorityQueue) push(e Element) {
	pq.Elements = append(pq.Elements, e)
	pq.sort()
}

func main() {
	m := readInput("input.txt")
	m.shortestPath([2]int{0, 0})
}

func (m *Mat) shortestPath(src [2]int) int {
	dest := [2]int{len(m.Elements) - 1, len(m.Elements[0]) - 1}
	pq := PriorityQueue{}
	pq.push(Element{Position: src, Priority: 0})
	totalRisk := map[[2]int]int{src:0}
	for {
		pos := pq.pop().Position
		if pos == dest {
			break
		}
		for _, n := range m.findAdjacentCells(pos) {
			riskSum := totalRisk[pos] + m.Elements[n[0]][n[1]]
			risk, ok := totalRisk[n]
			if !ok || riskSum < risk {
				totalRisk[n] = riskSum
				pq.push(Element{Position: n, Priority: m.manhattanDistance(n, dest)})
			}
		}
	}
	fmt.Println(totalRisk[dest])
	return totalRisk[dest]
}

func (m *Mat) findAdjacentCells(pos [2]int) [][2]int {
	adjacents := [][2]int{}
	adjacents = append(adjacents, [2]int{utils.MinMaxofInts(0, pos[0]-1, utils.MAX), pos[1]})
	adjacents = append(adjacents, [2]int{pos[0], utils.MinMaxofInts(0, pos[1]-1, utils.MAX)})
	adjacents = append(adjacents, [2]int{utils.MinMaxofInts(pos[0]+1, len(m.Elements)-1, utils.MIN), pos[1]})
	adjacents = append(adjacents, [2]int{pos[0], utils.MinMaxofInts(pos[1]+1, len(m.Elements)-1, utils.MIN)})
	return adjacents
}

func (m *Mat) manhattanDistance(src, dest [2]int) int {
	return int(math.Abs(float64(dest[0]-src[0])) + math.Abs(float64(dest[1]-src[1])))
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
		m.Elements = append(m.Elements, nums)
	}
	return m
}
