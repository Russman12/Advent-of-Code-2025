package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const inputFile = "./input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	start := time.Now()

	data, err := os.ReadFile(inputFile)
	check(err)

	width := strings.Index(string(data), "\n")

	str := strings.ReplaceAll(string(data), "\n", "")

	aMap := BuildAdjacency(str, width)

	_, removed := Peel4Core(aMap)

	fmt.Println(time.Since(start))
	fmt.Println(len(removed))
}

// BuildAdjacency builds adjacency for every '1' cell in bitstring, given row width W.
func BuildAdjacency(str string, W int) map[int][]int {
	n := len(str)
	if n%W != 0 {
		panic("bitstring length is not divisible by row width")
	}

	H := n / W
	adj := make(map[int][]int)

	// Holds the node ID (index) or -1 for each column of the previous/current row
	prevRow := make([]int, W)
	currRow := make([]int, W)
	for i := range prevRow {
		prevRow[i] = -1
		currRow[i] = -1
	}

	connect := func(a, b int) {
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	for r := range H {
		for c := range W {
			i := r*W + c
			if str[i] == '@' {
				currRow[c] = i
				if _, ok := adj[i]; !ok {
					adj[i] = []int{}
				}

				// W (left)
				if c > 0 && currRow[c-1] != -1 {
					connect(i, currRow[c-1])
				}
				// N (up)
				if prevRow[c] != -1 {
					connect(i, prevRow[c])
				}
				// NW (up-left)
				if c > 0 && prevRow[c-1] != -1 {
					connect(i, prevRow[c-1])
				}
				// NE (up-right)
				if c < W-1 && prevRow[c+1] != -1 {
					connect(i, prevRow[c+1])
				}
			} else {
				currRow[c] = -1
			}
		}

		// Shift current row to prev for next iteration
		copy(prevRow, currRow)
		for i := range currRow {
			currRow[i] = -1
		}
	}

	return adj
}

// Peel4Core takes an adjacency list and removes nodes with degree < 4
// until stable. It returns:
//   - remaining: nodes in the 4-core
//   - removed: nodes removed during peeling (in order)
func Peel4Core(adj map[int][]int) (remaining []int, removed []int) {
	// Current degree of each node
	degree := make(map[int]int, len(adj))
	for node, neighbors := range adj {
		degree[node] = len(neighbors)
	}

	// Queue of nodes that currently have degree < 4
	queue := make([]int, 0)
	inQueue := make(map[int]bool)

	for node, d := range degree {
		if d < 4 {
			queue = append(queue, node)
			inQueue[node] = true
		}
	}

	// Set to track removed nodes
	removedSet := make(map[int]bool)

	// Process queue
	for len(queue) > 0 {
		// pop last element
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if removedSet[node] {
			continue
		}

		removedSet[node] = true
		removed = append(removed, node)

		// Decrease degree of neighbors
		for _, nbr := range adj[node] {
			if removedSet[nbr] {
				continue
			}
			degree[nbr]--

			// If neighbor now has degree < 4, enqueue it
			if degree[nbr] < 4 && !inQueue[nbr] {
				queue = append(queue, nbr)
				inQueue[nbr] = true
			}
		}
	}

	// Remaining nodes = not removed
	for node := range adj {
		if !removedSet[node] {
			remaining = append(remaining, node)
		}
	}

	return
}

type node struct {
	edges []node
}
