package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	PART_ONE = false
	PART_TWO = true
)

type Graph struct {
	edges map[string][]string
	paths [][]string
}

func main() {
	graph := readInput("input.txt")
	fmt.Println(graph.findNumOfPaths(PART_TWO))
}

func (g *Graph) findNumOfPaths(visitTwice bool) int {
	g.traversePath("start", []string{}, visitTwice)
	return len(g.paths)
}

func (g *Graph) traversePath(node string, path []string, visitTwice bool) {
	path = append(path, node)
	for _, n := range g.edges[node] {
		if !g.isSmallCave(n) || g.countNode(path, n) == 0 {
			g.traversePath(n, path, visitTwice)
		} else if visitTwice && (!g.isSmallCave(n) || g.countNode(path, n) == 1) {
			g.traversePath(n, path, false)
		}
	}
	if path[len(path)-1] == "end" {
		g.paths = append(g.paths, path)
	}
}

func (g *Graph) countNode(slice []string, s string) int {
	var count int
	for _, n := range slice {
		if n == s {
			count++
		}
	}
	return count
}

func (g *Graph) isSmallCave(node string) bool {
	return strings.ToLower(node) == node
}

func readInput(fname string) Graph {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := Graph{}
	graph.edges = make(map[string][]string)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "-")
		if row[0] == "start" {
			graph.edges[row[0]] = append(graph.edges[row[0]], row[1])
		} else if row[1] == "start" {
			graph.edges[row[1]] = append(graph.edges[row[1]], row[0])
		} else {
			graph.edges[row[0]], graph.edges[row[1]] = append(graph.edges[row[0]], row[1]), append(graph.edges[row[1]], row[0])
		}
	}
	delete(graph.edges, "end")
	return graph
}
