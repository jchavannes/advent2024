package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	contents := string(file)
	re := regexp.MustCompile(`(mul\(([0-9]{1,3}),([0-9]{1,3})\))|(don't\(\))|(do\(\))`)
	matches := re.FindAllStringSubmatch(contents, -1)
	var sum, sumOn int
	var on = true
	for _, match := range matches {
		if match[0] == "don't()" {
			on = false
			continue
		} else if match[0] == "do()" {
			on = true
			continue
		}
		num1 := match[2]
		num2 := match[3]
		int1, err := strconv.Atoi(num1)
		if err != nil {
			log.Printf("match: %#v\n", match)
			log.Fatalf("error converting: %v %v", num1, err)
		}
		int2, err := strconv.Atoi(num2)
		if err != nil {
			log.Fatalf("error converting: %v %v", num2, err)
		}
		result := int1 * int2
		sum += result
		if on {
			sumOn += result
		}
	}
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Sum when on: %d\n", sumOn)
}
