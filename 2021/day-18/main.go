package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
	Level int
}

func main() {
	roots := readInput("test.txt")
	fmt.Println(roots)
	n := reduce(roots)
	fmt.Println("afterrrrr", n)
	fmt.Println(n.magnitude())
}

func reduce(roots []*Node) *Node {
	final := roots[0]
	for _, tree := range roots[1:] {
		final = final.add(tree)
		final.updateLevel(0)
		fmt.Println("----", final)
		for {
			if final.canExplode() {
				final.pairExplosion()
			}
			if !final.canExplode() && !final.split(final) {
				break
			} else {
				final.split(final)
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
	return &Node{Value: -1, Left: n, Right: m, Level: 0}
}

func (n *Node) updateLevel(level int) {
	n.Level = level
	if n.Left != nil && n.Right != nil {
		n.Left.updateLevel(level + 1)
		n.Right.updateLevel(level + 1)
	}
}

func (n *Node) split(next *Node) bool {
	if next.Value > 10 {
		next.Left = &Node{Value: next.Value / 2, Level: next.Level + 1}
		next.Right = &Node{Value: (next.Value + 1) / 2, Level: next.Level + 1}
		next.Value = -1
		return true
	}
	if next.Left != nil && next.Right != nil {
		next.split(next.Left)
		next.split(next.Right)
	}
	return false
}

func (n *Node) pairExplosion() {
	if n.Level == 3 && n.Value == -1 {
		if n.Right.Value > 0 {
			n.Right.Value += n.Left.Right.Value
			n.Left.Value = 0
			n.Left.Left, n.Left.Right = nil, nil
		} else if n.Left.Value > 0 {
			fmt.Println("---", n)
			n.Left.Value += n.Right.Left.Value
			n.Right.Value = 0
			n.Right.Left, n.Right.Right = nil, nil
		}
	}
	if n.Left != nil || n.Right != nil {
		n.Left.pairExplosion()
		n.Right.pairExplosion()
	}
	if n.Value >= 0 {

	}
}

func (n *Node) canExplode() bool {
	if n.Level == 4 && n.Value == -1 {
		// pair := n.Left
		// if n.Right.Value == -1 {
		// 	pair = n.Right
		// }
		return true
	}
	if n.Left != nil || n.Right != nil {
		n.Left.canExplode()
		n.Right.canExplode()
	}
	return false
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
