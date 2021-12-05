package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type BinaryInput struct {
	Row []string
	Col []string
}

func main() {
	input, _ := readInput("input.txt")
	rates := input.rates()
	p, _ := overall(string(rates[0]), string(rates[1]))
	o := input.rating(true)[0]
	c := input.rating(false)[0]
	r, _ := overall(o, c)
	fmt.Printf("part one: %d\npart two: %d", p, r)
}

//gamma and epsilon rates
func (b BinaryInput) rates() [2][]byte {
	var r [2][]byte
	re := regexp.MustCompile("1")
	for _, s := range b.Col {
		ones := re.FindAllString(s, -1)
		if len(ones) > len(b.Col[0])/2 {
			r[0] = append(r[0], 49)
			r[1] = append(r[1], 48)
		} else {
			r[0] = append(r[0], 48)
			r[1] = append(r[1], 49)
		}
	}
	return r
}

func (a *BinaryInput) rating(findMostCommon bool) []string {
	b := *a
	//match binary
	matchInt := 1
	//max of pos is row length
	pos := 1
	var ind [][]int
	for {
		re := regexp.MustCompile(fmt.Sprintf("%d", matchInt))
		ind = re.FindAllStringIndex(b.Col[0], -1)
		if len(ind) == 1 || len(b.Row) == 1 {
			b.Row = b.Row[ind[0][0]:ind[0][1]]
			break
		}
		switch findMostCommon {
		case true:
			if len(ind) > 1 && len(ind) < len(b.Row)/2 {
				re := regexp.MustCompile(fmt.Sprintf("%d", 1-matchInt))
				ind = re.FindAllStringIndex(b.Col[0], -1)
			} else if len(ind) == len(b.Row)/2 && matchInt != 1 {
				re := regexp.MustCompile("1")
				ind = re.FindAllStringIndex(b.Col[0], -1)
			}
		case false:
			if len(ind) > len(b.Row)/2 {
				re := regexp.MustCompile(fmt.Sprintf("%d", 1-matchInt))
				ind = re.FindAllStringIndex(b.Col[0], -1)
			} else if len(ind) == len(b.Row)/2 && matchInt != 0 {
				re := regexp.MustCompile("0")
				ind = re.FindAllStringIndex(b.Col[0], -1)
			}
		}
		var newRow []string
		for _, subInd := range ind {
			newRow = append(newRow, b.Row[subInd[0]:subInd[1]][0])
		}
		b.Row = newRow
		b.Col = b.Col[1:]
		b.Col[0] = updateNextCol(b.Row, pos)
		pos++
	}
	return b.Row
}

func overall(o string, c string) (int64, error) {
	a, err := strconv.ParseInt(o, 2, 64)
	if err != nil {
		return -1, err
	}
	b, err := strconv.ParseInt(c, 2, 64)
	if err != nil {
		return -1, err
	}
	return a * b, nil
}

func updateNextCol(rows []string, ind int) string {
	var col string
	for _, r := range rows {
		row := strings.Split(r, "")
		col += row[ind]
	}
	return col
}

func readInput(fname string) (input BinaryInput, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return BinaryInput{}, err
	}
	lines := strings.Split(string(b), "\n")
	col := make([]string, len(lines[0]))
	row := make([]string, len(lines))
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		row[i] = strings.TrimSpace(l)
		row := strings.Split(l, "")
		for j, n := range row {
			col[j] += n
		}
	}
	return BinaryInput{Col: col, Row: row}, nil
}
