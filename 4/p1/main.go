package main

import (
	"fmt"
	"os"
	"strings"
)

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

// very inefficient, but it works
func main() {
	input, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	var nonEmptyLines []string
	for _, line := range lines {
		if line != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}
	lines = nonEmptyLines

	if len(lines) == 0 {
		fmt.Println(0)
		return
	}

	lineLength := len(lines[0])
	for _, line := range lines {
		if len(line) != lineLength {
			fmt.Println("Error: lines have different lengths")
			return
		}
	}
	var count int
	// horizontal run
	for _, l := range lines {
		count += strings.Count(l, "XMAS")
		lr := reverse(l)
		count += strings.Count(lr, "XMAS")
	}
	// vertical run
	for i := 0; i < len(lines[0]); i++ {
		var s string
		for j := 0; j < len(lines); j++ {
			s += string(lines[j][i])
		}
		count += strings.Count(s, "XMAS")
		sr := reverse(s)
		count += strings.Count(sr, "XMAS")
	}
	// diagonal run
	for i := 0; i < len(lines); i++ {
		var s string
		for j := 0; j < len(lines)-i; j++ {
			s += string(lines[j][j+i])
		}
		count += strings.Count(s, "XMAS")
		sr := reverse(s)
		count += strings.Count(sr, "XMAS")
	}
	fmt.Println(count)
}
