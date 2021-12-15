package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/advent-of-code/2021/utils"
)

type Mat struct {
	Elements      [][]int
	FlashingCells [][]int
}

func main() {
	mat := readInput("input.txt")
	fmt.Printf("part one: %d\npart two: %d", mat.totalFlashes(100), mat.findFirstSync())
}

func (m *Mat) findFirstSync() int {
	var count int
	for {
		m.incrementByOne()
		for {
			if len(m.FlashingCells) == 0 {
				break
			}
			m.FlashingCells = m.incrementAdjacent(m.FlashingCells[0][0], m.FlashingCells[0][1])
		}
		count++
		if m.sum() == 0 {
			break
		}
	}
	return count
}

func (m *Mat) totalFlashes(steps int) int {
	var count int
	for range make([]struct{}, steps) {
		m.incrementByOne()
		for {
			if len(m.FlashingCells) == 0 {
				break
			}
			m.FlashingCells = m.incrementAdjacent(m.FlashingCells[0][0], m.FlashingCells[0][1])
		}
		count += m.countflashing()
	}
	return count
}

func (m *Mat) sum() (total int) {
	for _, row := range m.Elements {
		for _, num := range row {
			total += num
		}
	}
	return total
}

func (m *Mat) countflashing() (count int) {
	for _, row := range m.Elements {
		for _, num := range row {
			if num == 0 {
				count++
			}
		}
	}
	return count
}

func (m *Mat) incrementByOne() {
	for i := range m.Elements {
		for j := range m.Elements[i] {
			m.Elements[i][j] = (m.Elements[i][j] + 1) % 10
			if m.Elements[i][j] == 0 {
				m.FlashingCells = append(m.FlashingCells, []int{i, j})
			}
		}
	}
}

func (m *Mat) incrementAdjacent(row, col int) [][]int {
	rowLimit, colLimit := len(m.Elements)-1, len(m.Elements[0])-1
	for x := utils.minMaxofInts(0, row-1, utils.MIN); x <= utils.minMaxofInts(row+1, rowLimit, utils.MAX); x++ {
		for y := utils.minMaxofInts(0, col-1, utils.MIN); y <= utils.minMaxofInts(col+1, colLimit, utils.MAX); y++ {
			//if adjacent cell is already flashing, do not increment further
			if (x != row || y != col) && m.Elements[x][y] != 0 {
				m.Elements[x][y] = (m.Elements[x][y] + 1) % 10
				if m.Elements[x][y] == 0 {
					m.FlashingCells = append(m.FlashingCells, []int{x, y})
				}
			}
		}
	}
	return m.FlashingCells[1:]
}

func readInput(fname string) Mat {
	mat := Mat{}
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		rowNums := make([]int, 0, len(row))
		for _, r := range row {
			num, err := strconv.Atoi(r)
			if err != nil {
				log.Fatal(err)
			}
			rowNums = append(rowNums, num)
		}
		mat.Elements = append(mat.Elements, rowNums)
	}
	return mat
}
