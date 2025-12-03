package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

const digitCnt = 12

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile(inputFile)
	check(err)

	dGroups := strings.Split(string(data), "\n")
	sum := 0
	for _, val := range dGroups {
		digits := strings.Split(val, "")

		if len(digits) > 0 {
			sum += getMaxVal(digits)
		}
	}

	fmt.Println(sum)
}

func getMaxVal(digits []string) int {
	strs := []string{}
	currentIdx := -1

	for i := 0; i < digitCnt; i++ {
		maxVal := 0
		endIdx := len(digits) - (digitCnt - 1 - i)
		maxVal, currentIdx = maxValMinMax(digits, currentIdx+1, endIdx)
		strs = append(strs, strconv.Itoa(maxVal))
	}

	val, err := strconv.Atoi(strings.Join(strs, ""))
	check(err)
	return val
}

func maxValMinMax(digits []string, startIdx int, endIdx int) (int, int) {
	maxVal, maxIdx := 0, 0
	for i := startIdx; i < endIdx; i++ {

		val, err := strconv.Atoi(digits[i])
		check(err)

		if val > maxVal {
			maxVal = val
			maxIdx = i
		}
	}

	return maxVal, maxIdx
}
