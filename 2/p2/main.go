package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func simpleSafe(line []int) bool {
	var dec bool
	for i := 0; i < len(line)-1; i++ {
		res := line[i] - line[i+1]
		if i == 0 {
			if res < 0 {
				dec = true // dec is false by default
			}
		}
		if res == 0 {
			return false
		}
		if dec && res > 0 || !dec && res < 0 {
			return false
		}
		if abs(res) > 3 {
			return false
		}
	}
	return true
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

// PROBLEM: does not correctly determine insafeIndex when dec/asc causes unsafe; need to redo dec logic
func safe(line []int) bool {
	var dec bool
	var unsafeIndexes []int
	for i := 0; i < len(line)-1; i++ {
		res := line[i] - line[i+1]
		if i == 0 {
			if res < 0 {
				dec = true // dec is false by default
			}
		}
		if res == 0 {
			unsafeIndexes = append(unsafeIndexes, i)
			continue
		}
		if dec && res > 0 || !dec && res < 0 {
			unsafeIndexes = append(unsafeIndexes, i)
			continue
		}
		if abs(res) > 3 {
			unsafeIndexes = append(unsafeIndexes, i)
			continue
		}
	}
	if len(unsafeIndexes) == 0 {
		return true
	}
	for _, i := range unsafeIndexes {
		t := make([]int, len(line))
		copy(t, line)
		t = remove(t, i)
		if simpleSafe(t) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := bufio.NewScanner(file)
	var safeNum int
	for lines.Scan() {
		s := strings.Split(lines.Text(), " ")
		line := make([]int, len(s))
		for i, x := range s {
			y, _ := strconv.Atoi(x)
			line[i] = y
		}
		if safe(line) {
			safeNum++
			continue
		}

	}
	fmt.Println(safeNum)
}
