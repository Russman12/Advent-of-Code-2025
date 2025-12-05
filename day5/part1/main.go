package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

const inputFile = "./input.txt"

var wg sync.WaitGroup

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

	ranges := [][]int{}
	for _, rangeStr := range rangeStrs {
		rangeArr := strings.Split(rangeStr, "-")
		lb, err := strconv.Atoi(rangeArr[0])
		check(err)
		ub, err := strconv.Atoi(rangeArr[1])
		check(err)
		ranges = append(ranges, []int{lb, ub})
	}

	idStrs := strings.Split(parts[1], "\n")
	ids := []int{}
	for _, str := range idStrs {
		if str == "" {
			continue
		}
		id, err := strconv.Atoi(str)
		check(err)

		ids = append(ids, id)
	}

	// fmt.Println(ids)

	cntChan := make(chan int)

	var wg2 sync.WaitGroup
	wg2.Add(1)
	freshIds := []int{}
	go func() {
		defer wg2.Done()
		for id := range cntChan {
			if !slices.Contains(freshIds, id) {
				freshIds = append(freshIds, id)
			}
		}
	}()

	for _, r := range ranges {
		wg.Add(1)
		go evalRange(r[0], r[1], ids, cntChan)
	}

	wg.Wait()
	close(cntChan)
	wg2.Wait()

	fmt.Println(time.Since(start))
	fmt.Println(len(freshIds))
}

func evalRange(lowerBound, upperBound int, ids []int, c chan int) {
	defer wg.Done()
	// fmt.Println(lowerBound, upperBound)
	for _, id := range ids {
		if id >= lowerBound && id <= upperBound {
			c <- id
		}
	}
}
