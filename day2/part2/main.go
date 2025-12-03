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
	divisors := []int{}
	for i := 1; i <= len(id)/2; i++ {
		if len(id)%i == 0 {
			divisors = append(divisors, i)
		}
	}

	for _, divisor := range divisors {
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
