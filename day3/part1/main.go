package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

const digitCnt = 2

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

		sum += getMaxVal(digits)
	}

	fmt.Println(sum)
}

func getMaxVal(digits []string) int {
	maxVal, maxIdx := 0, 0
	for i := 0; i < len(digits)-1; i++ {
		val, err := strconv.Atoi(digits[i])
		check(err)

		if val > maxVal {
			maxVal = val
			maxIdx = i
		}
	}
	secondMax := 0
	for i := maxIdx + 1; i < len(digits); i++ {
		val, err := strconv.Atoi(digits[i])
		check(err)

		if val > secondMax {
			secondMax = val
		}
	}

	return maxVal*10 + secondMax
}
