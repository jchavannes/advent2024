package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("locations.txt")
	if err != nil {
		panic(err)
	}
	contents := string(file)
	var lines = strings.Split(contents, "\n")
	var leftNums = make([]int, len(lines))
	var rightNums = make([]int, len(lines))
	for i, line := range lines {
		nums := strings.Fields(line)
		leftNums[i], err = strconv.Atoi(strings.TrimSpace(nums[0]))
		if err != nil {
			log.Fatalf("error converting: %v %v", nums, err)
		}
		rightNums[i], err = strconv.Atoi(strings.TrimSpace(nums[1]))
		if err != nil {
			log.Fatalf("error converting: %v %v", nums, err)
		}
	}
	sort.Slice(leftNums, func(i, j int) bool { return leftNums[i] < leftNums[j] })
	sort.Slice(rightNums, func(i, j int) bool { return rightNums[i] < rightNums[j] })
	var totalDistance int
	for i := range leftNums {
		distance := rightNums[i] - leftNums[i]
		if distance < 0 {
			distance = -distance
		}
		totalDistance += distance
	}
	fmt.Printf("Total distance: %d\n", totalDistance)
	var totalSimilarity int
	for _, leftNum := range leftNums {
		var found int
		for _, rightInt := range rightNums {
			if leftNum == rightInt {
				found++
			}
		}
		var similarity = leftNum * found
		totalSimilarity += similarity
	}
	fmt.Printf("Total similarity: %d\n", totalSimilarity)
}
