package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Mat struct {
	Elements [][]int
	LocalMin [][2]int
}

func main() {
	mat := readInput("input.txt")
	fmt.Printf("part one: %d\npart two: %d", mat.calcRiskLevel(), mat.multiplyLargestBasin())
}

func (m *Mat) multiplyLargestBasin()int{
	var basinSize []int
	for _, ind := range m.LocalMin {
		basinSize = append(basinSize, m.basinSizeForLocalMin(ind[0], ind[1]))
	}
	sort.Ints(basinSize) 
	res := 1
	for _,num := range basinSize[len(basinSize)-3:] {
		res *= num
	}
	return res
}

func (m *Mat) calcRiskLevel() (risk int) {
	for i := 1; i < len(m.Elements)-1; i++ {
		for j := 1; j < len(m.Elements[i])-1; j++ {
			if m.isLocalMin(i, j) {
				risk += (m.Elements[i][j] + 1)
				m.LocalMin = append(m.LocalMin, [2]int{i, j})
			}
		}
	}
	return risk
}

func (m *Mat) basinSizeForLocalMin(row, col int) int {
	if m.Elements[row][col] == 9 {
		return 0
	}
	m.Elements[row][col] = 9
	return 1 + m.basinSizeForLocalMin(row-1, col) + m.basinSizeForLocalMin(row, col-1) + m.basinSizeForLocalMin(row+1, col) + m.basinSizeForLocalMin(row, col+1)
}

func (m *Mat) isLocalMin(row, col int) bool {
	if m.Elements[row][col] == 9 {
		return false
	}
	min := m.Elements[row][col]
	for dxy := -1; dxy <= 1; dxy = dxy + 2 {
		if m.Elements[row][col+dxy] < min {
			min = m.Elements[row][col+dxy]
		}
		if m.Elements[row+dxy][col] < min {
			min = m.Elements[row+dxy][col]
		}
	}
	return min == m.Elements[row][col]
}

func readInput(fname string) Mat {
	mat := Mat{}
	file, err := os.Open(fname)
	if err != nil {

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	makeRowPadding := true
	for scanner.Scan() {
		row := strings.Split(fmt.Sprintf("%d%v%d", 9, scanner.Text(), 9), "")
		rowNums := make([]int, 0, len(row))
		if makeRowPadding == true {
			mat.Elements = append(mat.Elements, makePadding(len(row)))
			makeRowPadding = false
		}
		for _, r := range row {
			num, err := strconv.Atoi(r)
			if err != nil {
				log.Fatal(err)
			}
			rowNums = append(rowNums, num)
		}
		mat.Elements = append(mat.Elements, rowNums)
	}
	mat.Elements = append(mat.Elements, makePadding(len(mat.Elements[0])))
	return mat
}

func makePadding(length int) []int {
	nums := make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = 9
	}
	return nums
}
