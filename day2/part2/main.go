package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const inputFile = "./input.txt"

var lenDivisors = map[int][]int{
	2:  {1},
	3:  {1},
	4:  {1, 2},
	5:  {1},
	6:  {1, 2, 3},
	7:  {1},
	8:  {1, 2, 4},
	9:  {1, 3},
	10: {1, 2, 5},
	11: {1},
	12: {1, 2, 3, 4, 6},
	13: {1},
	14: {1, 2, 7},
	15: {1, 3, 5},
	16: {1, 2, 4, 8},
	17: {1},
	18: {1, 2, 3, 6, 9},
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(inputFile)
	check(err)

	strs := strings.Split(string(data), ",")

	totalSum := 0
	for _, val := range strs {
		lbubSlice := strings.Split(strings.Trim(val, "\n"), "-")

		lowerBound, err := strconv.Atoi(lbubSlice[0])
		check(err)

		upperBound, err := strconv.Atoi(lbubSlice[1])
		check(err)

		totalSum += rangeInvalidSum(lowerBound, upperBound)
	}

	fmt.Println(totalSum)

	fmt.Println(time.Since(start))
}

func rangeInvalidSum(lowerBound int, upperBound int) int {
	sum := 0

	for i := lowerBound; i <= upperBound; i++ {
		if !isValid(strconv.Itoa(i)) {
			sum += i
		}
	}
	return sum
}

func isValid(id string) bool {
	for _, divisor := range lenDivisors[len(id)] {
		if allEqual(strSplitN(id, divisor)) {
			return false
		}
	}

	return true
}

func strSplitN(str string, n int) []string {
	strs := []string{}
	for i := 0; i < len(str); i += n {
		strs = append(strs, str[i:i+n])
	}
	return strs
}

func allEqual(strs []string) bool {
	first := strs[0]

	for i := 1; i < len(strs); i++ {
		if first != strs[i] {
			return false
		}
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
