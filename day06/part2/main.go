package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
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

	fns := loadData()

	solution := 0
	for _, fn := range fns {
		solution += fn.calc()
	}

	fmt.Println(time.Since(start))
	fmt.Println(solution)
}

func loadData() (fns []function) {
	data, err := os.ReadFile(inputFile)
	check(err)

	lines := bytes.Split(bytes.Trim(data, "\n"), []byte("\n"))

	for i, char := range bytes.Trim(lines[len(lines)-1], "\n") {
		if char != byte(' ') {
			fns = append(fns, function{i, char, [][]byte{}})
		}
	}

	for j, fn := range fns {
		inputs := [][]byte{}
		endIdx := 0
		if j < len(fns)-1 {
			endIdx = fns[j+1].startIdx - 1
		} else {
			endIdx = -1
		}

		for i := 0; i < len(lines)-1; i++ {
			if endIdx == -1 {
				inputs = append(inputs, lines[i][fn.startIdx:])
			} else {
				inputs = append(inputs, lines[i][fn.startIdx:endIdx])
			}
		}
		fns[j].inputs = inputs
	}

	return
}

type function struct {
	startIdx int
	operator byte
	inputs   [][]byte
}

func (f *function) vals() (vals []int) {
	valsBytes := [][]byte{}
	for _, bts := range f.inputs {
		for j, bt := range bts {
			if len(valsBytes)-1 < j {
				valsBytes = append(valsBytes, []byte{})
			}

			valsBytes[j] = append(valsBytes[j], bt)
		}
	}

	for _, valBytes := range valsBytes {
		val, err := strconv.Atoi(string(bytes.Trim(valBytes, " ")))
		check(err)

		vals = append(vals, val)
	}

	return
}

func (f *function) calc() int {
	vals := f.vals()
	var calc int
	if f.operator == '*' {
		calc = 1
		for _, val := range vals {
			calc *= val
		}
	}
	if f.operator == '+' {
		calc = 0
		for _, val := range vals {
			calc += val
		}
	}

	return calc
}
