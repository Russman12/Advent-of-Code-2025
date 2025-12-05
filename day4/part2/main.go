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

	neighbors, deg, active := BuildAdjacency(str, width)

	removed, _ := Peel4Core(neighbors, deg, active)

	fmt.Println(time.Since(start))
	fmt.Println(len(removed))
}

func BuildAdjacency(str string, W int) (neighbors [][8]int, deg []int, active []bool) {
	n := len(str)
	H := n / W

	neighbors = make([][8]int, n)
	deg = make([]int, n)
	active = make([]bool, n)

	// fill with -1 (no neighbor)
	for i := range neighbors {
		for k := range 8 {
			neighbors[i][k] = -1
		}
	}

	// Previous and current row of node IDs or -1
	prev := make([]int, W)
	curr := make([]int, W)
	for i := range prev {
		prev[i] = -1
		curr[i] = -1
	}

	add := func(a, b int) {
		// add b to a
		for k := range 8 {
			if neighbors[a][k] == -1 {
				neighbors[a][k] = b
				deg[a]++
				return
			}
			if neighbors[a][k] == b {
				// already present (defensive), don't double-count
				return
			}
		}
	}

	for r := range H {
		for c := range W {
			idx := r*W + c
			if str[idx] == '@' {
				active[idx] = true
				curr[c] = idx

				// check 4 possible neighbors
				if c > 0 && curr[c-1] != -1 { // left
					add(idx, curr[c-1])
					add(curr[c-1], idx)
				}
				if prev[c] != -1 { // up
					add(idx, prev[c])
					add(prev[c], idx)
				}
				if c > 0 && prev[c-1] != -1 { // up-left
					add(idx, prev[c-1])
					add(prev[c-1], idx)
				}
				if c < W-1 && prev[c+1] != -1 { // up-right
					add(idx, prev[c+1])
					add(prev[c+1], idx)
				}

			} else {
				curr[c] = -1
			}
		}
		copy(prev, curr)
		for i := range curr {
			curr[i] = -1
		}
	}

	return neighbors, deg, active
}

func Peel4Core(neighbors [][8]int, deg []int, active []bool) (removed []int, remains []int) {
	n := len(neighbors)

	queue := make([]int, 0, n)
	inQueue := make([]bool, n)
	removedFlag := make([]bool, n)

	// enqueue all active nodes with degree < 4 (including deg == 0)
	for i := range n {
		if !active[i] {
			continue
		}
		if deg[i] < 4 { // <-- include deg == 0
			queue = append(queue, i)
			inQueue[i] = true
		}
	}

	for len(queue) > 0 {
		v := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if removedFlag[v] {
			continue
		}

		removedFlag[v] = true
		removed = append(removed, v)

		for _, nb := range neighbors[v] {
			if nb == -1 {
				continue
			}
			if removedFlag[nb] {
				continue
			}

			deg[nb]--
			if deg[nb] < 4 && !inQueue[nb] {
				queue = append(queue, nb)
				inQueue[nb] = true
			}
		}
	}
	// collect remains only among active nodes
	for i := range n {
		if !active[i] {
			continue
		}
		if deg[i] >= 4 {
			remains = append(remains, i)
		}
	}
	return
}
