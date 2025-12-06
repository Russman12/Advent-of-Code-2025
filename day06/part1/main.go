package main

import (
	"fmt"
	"os"
	"strconv"
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

	ops, vals := loadData()
	solution := solve(ops, vals)

	fmt.Println(time.Since(start))
	fmt.Println(solution)
}

func solve(ops []string, vals [][]string) int {
	solution := 0
	for i, op := range ops {
		var calc int
		if op == "*" {
			calc = 1
			for _, val := range vals {
				m, err := strconv.Atoi(val[i])
				check(err)

				calc *= m
			}
		}
		if op == "+" {
			calc = 0
			for _, val := range vals {
				m, err := strconv.Atoi(val[i])
				check(err)

				calc += m
			}
		}
		solution += calc
	}

	return solution
}

func loadData() (ops []string, vals [][]string) {
	data, err := os.ReadFile(inputFile)
	check(err)

	lines := strings.Split(strings.Trim(string(data), "\n"), "\n")

	ops = strings.Fields(lines[len(lines)-1])

	for i := 0; i < len(lines)-1; i++ {
		vals = append(vals, strings.Fields(lines[i]))
	}

	return
}
