package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/dmytzo/adventofcode2024/tool"
)

const separator = "   "

func main() {
	var leftIDs, rightIDs []int
	rightIDsMap := map[int]int{}

	for _, line := range tool.InputLines("day01/input.txt") {
		ids := strings.Split(line, separator)
		if len(ids) != 2 {
			log.Fatalf("invalid line: %s\n", line)
		}

		leftID := tool.MustInt(ids[0])
		leftIDs = append(leftIDs, leftID)

		rightID := tool.MustInt(ids[1])
		rightIDs = append(rightIDs, rightID)
		rightIDsMap[rightID]++
	}

	slices.Sort(leftIDs)
	slices.Sort(rightIDs)

	if len(leftIDs) != len(rightIDs) {
		log.Fatalf("lines num is not equal")
	}

	var task1Result int
	var task2Result int

	for idx, v := range leftIDs {
		task1Result += int(math.Abs(float64(v - rightIDs[idx])))
		task2Result += v * rightIDsMap[v]
	}

	fmt.Printf(`
# RESULTS
Task 1: %d
Task 2: %d
`, task1Result, task2Result)
}
