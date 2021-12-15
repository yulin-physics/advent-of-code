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

type Mat struct {
	Elements [][]int
}

func main() {
	m := readInput(100, 500, "input.txt")
	m.expand()
	fmt.Printf("the risk level of lowest risk path for a %dx%d cavern: %d\n", len(m.Elements), len(m.Elements[0]), m.shortestWeightedPath([2]int{0, 0}))
}

func (m *Mat) expand() {
	originalLen := len(m.Elements)
	m.Elements = m.Elements[:cap(m.Elements)]
	for i := range make([]struct{}, originalLen) {
		m.Elements[i] = m.Elements[i][:cap(m.Elements[i])]
		for j := originalLen; j < len(m.Elements); j++ {
			new := (m.Elements[i][j-originalLen] + 1) % 10
			if new == 0 {
				m.Elements[i][j] = 1
			} else {
				m.Elements[i][j] = new
			}
		}
	}
	for i := originalLen; i < len(m.Elements); i++ {
		m.Elements[i] = make([]int, len(m.Elements))
		for j := range make([]struct{}, len(m.Elements)) {
			new := (m.Elements[i-originalLen][j] + 1) % 10
			if new == 0 {
				m.Elements[i][j] = 1
			} else {
				m.Elements[i][j] = new
			}
		}
	}
}

func (m *Mat) shortestWeightedPath(src [2]int) int {
	dest := [2]int{len(m.Elements) - 1, len(m.Elements[0]) - 1}
	pq := PriorityQueue{}
	pq.push(Element{Position: src, Priority: 0})
	totalRisk := map[[2]int]int{src: 0}
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
				//set priority metric as furthest distance from dest, dest at the back of the queue
				pq.push(Element{Position: n, Priority: m.manhattanDistance(n, dest)})
			}
		}
	}
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

type Element struct {
	Position [2]int
	Priority int
}

type PriorityQueue struct {
	Elements []Element
}

func (pq *PriorityQueue) pop() Element {
	first := pq.Elements[0]
	pq.Elements = pq.Elements[1:]
	return first
}

func (pq *PriorityQueue) push(e Element) {
	pq.Elements = append(pq.Elements, e)
	sort.SliceStable(pq.Elements, func(i, j int) bool {
		return pq.Elements[i].Priority > pq.Elements[j].Priority
	})
}

func readInput(oneTileSize, fullSize int, fname string) Mat {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := Mat{}
	m.Elements = make([][]int, oneTileSize, fullSize)
	i := 0
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		m.Elements[i] = make([]int, oneTileSize, fullSize)
		for j, r := range row {
			num, err := strconv.Atoi(r)
			if err != nil {
				log.Fatal(err)
			}
			m.Elements[i][j] = num
		}
		i++
	}
	return m
}
