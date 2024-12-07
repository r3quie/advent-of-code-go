package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	input, _ := os.ReadFile("input.txt")

	pattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)

	mulFlag := true
	var r int
	ddPattern := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	dd := ddPattern.FindAllString(string(input), -1)

	for _, x := range dd {
		if doPattern.MatchString(x) {
			mulFlag = true
		} else if dontPattern.MatchString(x) {
			mulFlag = false
		} else if mulFlag {
			match := pattern.FindStringSubmatch(x)
			if match != nil {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				r += a * b
			}
		}
	}
	fmt.Println(r)
}
