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
	Value   int
	Left    *Node
	Right   *Node
	Level   int
	Visited bool
}

func main() {
	roots := readInput("test.txt")
	reduce(roots)
	// n := reduce(roots)

	// for _, n := range roots {
	// 	ok, node := n.canExplode()
	// 	if ok {
	// 		n.pairExplosion(node, n)
	// 	}
	// 	fmt.Println(n)
	// }
	// fmt.Println("afterrrrr", n)
	// fmt.Println(n.magnitude())
}

func reduce(roots []*Node) *Node {
	final := roots[0]
	fmt.Println(roots[1])

	for _, tree := range roots[1:] {
		final = final.add(tree)
		final.updateLevel(0)
		fmt.Println("----", final)
		// for {
		ok, pair := final.canExplode()
		if ok {
			final.pairExplosion(pair, final)
		}
		//
		fmt.Println("**", final)
		ok, pair = final.canExplode()
		if ok {
			final.pairExplosion(pair, final)
		}
		//
		fmt.Println("**", final)
		final.split()
		fmt.Println("**", final)

	}

	// for {
	// 	// if final.canExplode() {
	// 	// 	final.pairExplosion()
	// 	// }
	// 	// if !final.canExplode() && !final.split(final) {
	// 	// 	break
	// 	// } else {
	// 	// 	final.split(final)
	// 	// }

	// }

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
		return n.Right.split()
	}
	return left
}

func (n *Node) checkRegNumLeft(pair *Node) {
	if n.Right.Value >= 0 {
		fmt.Println("000000000", n.Right)
		n.Right.Value += pair.Left.Value
		return
	}
	n.Left.checkRegNumLeft(pair)
}

func (n *Node) checkRegNumRight(pair *Node) {
	if n.Left.Value >= 0 {
		n.Left.Value += pair.Right.Value
		return
	}
	n.Right.checkRegNumRight(pair)
}

func (n *Node) pairExplosion(pair *Node, parent *Node) bool {
	if n == nil {
		return false
	}
	if n.Level == 3 && n.Value == -1 {
		if reflect.DeepEqual(n.Right, pair) {
			parent.checkRegNumLeft(pair)
			n.Left.Value += pair.Left.Value
			n.Right.Value = 0
			n.Right.Left, n.Right.Right = nil, nil
			return true
		} else if reflect.DeepEqual(n.Left, pair) {
			parent.checkRegNumRight(pair)
			n.Right.Value += pair.Right.Value
			n.Left.Value = 0
			n.Left.Left, n.Left.Right = nil, nil
			return true
		}
	}
	left := n.Left.pairExplosion(pair, parent)
	if !left {
		return n.Right.pairExplosion(pair, parent)
	}
	return left
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
	// ok, left := n.Left.canExplode()
	// if ok {
	// 	return true, left
	// }
	// return n.Right.canExplode()
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
