package main

import (
	"fmt"
	"os"
	"sort"
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

	data, err := os.ReadFile(inputFile)
	check(err)

	parts := strings.Split(string(data), "\n\n")

	rangeStrs := strings.Split(parts[0], "\n")

	ranges := [][2]int{}
	for _, rangeStr := range rangeStrs {
		rangeArr := strings.Split(rangeStr, "-")
		lb, err := strconv.Atoi(rangeArr[0])
		check(err)
		ub, err := strconv.Atoi(rangeArr[1])
		check(err)
		ranges = append(ranges, [2]int{lb, ub})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	solution := 0
	currentRange := ranges[0]
	for i := 1; i < len(ranges); i++ {
		// lowerbound < current upper bound and upper bound > current upper bound
		if ranges[i][0] <= currentRange[1] && ranges[i][1] > currentRange[1] {
			currentRange[1] = ranges[i][1]
		}

		if ranges[i][0] > currentRange[1] {
			solution += currentRange[1] - currentRange[0] + 1
			currentRange = ranges[i]
		}
	}

	solution += currentRange[1] - currentRange[0] + 1

	fmt.Println(time.Since(start))
	fmt.Println(solution)
}
