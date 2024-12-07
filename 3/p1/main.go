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
	matches := pattern.FindAllStringSubmatch(string(input), -1)

	var r int

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		r += x * y
	}

	fmt.Println(r)
}
