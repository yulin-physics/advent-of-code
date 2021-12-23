package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Node struct {
	Value               int
	Left, Right, Parent *Node
	Level               int
	Visited             bool
}

func main() {
	roots := readInput("test.txt")
	reduce(roots)
	n := reduce(roots)
	fmt.Println("afterrrrr", n)
	fmt.Println(n.magnitude())
}

func reduce(roots []*Node) *Node {
	final := roots[0]
	fmt.Println(roots[1])

	for _, tree := range roots[1:] {
		final = final.add(tree)
		final.updateLevel(0)
		for {
			okExplode, pair := final.canExplode()
			if okExplode {
				final.pairExplosion(pair, final, final)
				continue
			}
			okSplit := final.split()
			if !okExplode && !okSplit{
				break
			}
		}
	}

	return final
}

func (n *Node) magnitude() int {
	if n == nil {
		return 0
	}
	if n.Value >= 0 {
		return n.Value
	}
	return 3*n.Left.magnitude() + 2*n.Right.magnitude()
}

func (n *Node) add(m *Node) *Node {
	//TODO: call reduce
	return &Node{Value: -1, Left: n, Right: m, Level: 0}
}

func (n *Node) updateLevel(level int) {
	n.Level = level
	n.Visited = false
	if n.Left != nil && n.Right != nil {
		n.Left.updateLevel(level + 1)
		n.Right.updateLevel(level + 1)
	}
}

func (n *Node) split() bool {
	if n == nil {
		return false
	}
	if n.Value > 10 {
		n.Left = &Node{Value: n.Value / 2, Level: n.Level + 1}
		n.Right = &Node{Value: (n.Value + 1) / 2, Level: n.Level + 1}
		n.Value = -1
		return true
	}
	left := n.Left.split()
	if !left {
		left = n.Right.split()
	}
	return left
}

func (n *Node) pairExplosion(pair *Node, parent *Node, previous *Node)  {
	if n == nil {
		fmt.Println(n, "here")
		return 
	}
	if reflect.DeepEqual(n.Right, pair) {
		if previous.Right.Value > 0 {
			previous.Right.Value += pair.Right.Value
		}
		n.Left.Value += pair.Left.Value
		n.Right.Value = 0
		n.Right.Left, n.Right.Right = nil, nil
		return
	} else if reflect.DeepEqual(n.Left, pair) {
		if previous.Left.Value > 0 {
			previous.Left.Value += pair.Left.Value
		}
		n.Right.Value += pair.Right.Value
		n.Left.Value = 0
		n.Left.Left, n.Left.Right = nil, nil
		return 
	}
	n.Left.pairExplosion(pair, parent, n)
	n.Right.pairExplosion(pair, parent, n)
}

func (n *Node) canExplode() (bool, *Node) {
	ok := false
	found := &Node{}
	if n.Left != nil {
		ok, found = n.Left.canExplode()
	}
	if !ok && n.Right != nil {
		ok, found = n.Right.canExplode()
	}
	if n.Level == 4 && n.Value == -1 {
		return true, n
	}
	return ok, found
}

func (n *Node) String() string {
	return fmt.Sprintf("{%d, %v, %v, %d}", n.Value, n.Left, n.Right, n.Level)
}

func initTree(line string, level int) *Node {
	if !strings.HasPrefix(line, "[") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		return &Node{Value: val, Level: level}
	}
	count := 0
	for i, r := range line {
		switch r {
		case '[':
			count++
		case ']':
			count--
		case ',':
			if count == 1 {
				return &Node{Value: -1, Left: initTree(line[1:i], level+1), Right: initTree(line[i+1:len(line)-1], level+1), Level: level}
			}
		}
	}
	return nil
}

func readInput(fname string) []*Node {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	roots := []*Node{}
	for scanner.Scan() {
		node := initTree(scanner.Text(), 0)
		roots = append(roots, node)
	}
	return roots
}
