package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

func main() {
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
	if len(id)%2 != 0 {
		return true
	}

	return id[:len(id)/2] != id[len(id)/2:]
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
