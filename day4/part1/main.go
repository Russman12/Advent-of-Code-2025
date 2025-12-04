package main

import (
	"fmt"
	"os"
	"strings"
)

const inputFile = "./input-test.txt"

const (
	distance    = 1
	adjacentMax = 4
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile(inputFile)
	check(err)

	strs := strings.Split(strings.Trim(string(data), "\n"), "\n")

	grid := [][]string{}

	for _, str := range strs {
		grid = append(grid, strings.Split(str, ""))
	}

	cnt := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == "@" && canPick(grid, x, y) {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}

func canPick(grid [][]string, x int, y int) bool {
	cnt := 0
	for checkY := y - distance; checkY <= y+distance; checkY++ {
		if checkY < 0 || checkY == len(grid) {
			continue
		}
		for checkX := x - distance; checkX <= x+distance; checkX++ {
			if checkX < 0 || checkX == len(grid[checkY]) {
				continue
			}

			if checkY == y && checkX == x {
				continue
			}

			if grid[checkY][checkX] == "@" {
				cnt++
				if cnt >= adjacentMax {
					return false
				}
			}
		}
	}

	return true
}
