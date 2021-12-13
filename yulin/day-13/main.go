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

type Origami struct {
	horizontal   map[int][]int
	vertical     map[int][]int
	instructions [][2]int
}

func main() {
	origami := readInput("input.txt")
	origami.followInstructions(true)
	origami.draw()
}

func (o *Origami) draw() {
	keys := []int{}
	for k := range o.vertical {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		row := o.makeLine(50)
		for _, x := range o.vertical[k] {
			row[x] = '#'
		}
		fmt.Println(string(row))
	}
}

func (o *Origami) makeLine(length int) []rune {
	row := make([]rune, length)
	for i := range row {
		row[i] = '.'
	}
	return row
}

func (o *Origami) followInstructions(printPartOne bool) {
	for _, ins := range o.instructions {
		if ins[1] == 0 {
			for k, v := range o.horizontal {
				if k > ins[0] {
					o.horizontal[ins[0]-(k-ins[0])] = append(o.horizontal[ins[0]-(k-ins[0])], v...)
					delete(o.horizontal, k)
				}
			}
			o.updateYAxis()
		} else {
			for k, v := range o.vertical {
				if k > ins[1] {
					o.vertical[ins[1]-(k-ins[1])] = append(o.vertical[ins[1]-(k-ins[1])], v...)
					delete(o.vertical, k)
				}
			}
			o.updateXAxis()
		}
		if printPartOne {
			fmt.Printf("number of visible dots after first fold: %d\n", o.findVisibleDots())
			printPartOne = false
		}
	}
}

func (o *Origami) findVisibleDots() int {
	var sum int
	for k, v := range o.horizontal {
		o.horizontal[k] = o.removeDuplicateInt(v)
		sum += len(o.horizontal[k])
	}
	return sum
}

func (o *Origami) removeDuplicateInt(slice []int) []int {
	exist := make(map[int]struct{})
	newSlice := []int{}
	for _, num := range slice {
		if _, val := exist[num]; !val {
			exist[num] = struct{}{}
			newSlice = append(newSlice, num)
		}
	}
	return newSlice
}

func (o *Origami) updateYAxis() {
	new := make(map[int][]int)
	for k, v := range o.horizontal {
		for _, x := range v {
			new[x] = append(new[x], k)
		}
	}
	o.vertical = new
}

func (o *Origami) updateXAxis() {
	new := make(map[int][]int)
	for k, v := range o.vertical {
		for _, y := range v {
			new[y] = append(new[y], k)
		}
	}
	o.horizontal = new
}

func readInput(fname string) Origami {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	o := Origami{}
	o.horizontal, o.vertical = make(map[int][]int), make(map[int][]int)
	instructions := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			instructions = true
			continue
		}
		if instructions {
			row := strings.Split(strings.Replace(scanner.Text(), "fold along ", "", -1), "=")
			coord, _ := strconv.Atoi(row[1])
			if row[0] == "x" {
				o.instructions = append(o.instructions, [2]int{coord, 0})
			} else if row[0] == "y" {
				o.instructions = append(o.instructions, [2]int{0, coord})
			}
		} else {
			row := strings.Split(scanner.Text(), ",")
			x, _ := strconv.Atoi(row[0])
			y, _ := strconv.Atoi(row[1])
			o.horizontal[x] = append(o.horizontal[x], y)
			o.vertical[y] = append(o.vertical[y], x)
		}
	}
	return o
}
