// I did not try optimizing this code; this is a brute force solution.

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

func safe(line []int) bool {
	var dec bool
	var flag bool
	for i := 0; i < len(line)-1; i++ {
		res := line[i] - line[i+1]
		if i == 0 {
			if res < 0 {
				dec = true // dec is false by default
			}
		}
		if res == 0 {
			flag = true
			break
		}
		if dec && res > 0 || !dec && res < 0 {
			flag = true
			break
		}
		if abs(res) > 3 {
			flag = true
			break
		}
	}
	if !flag {
		return true
	}
	for i := range line {
		t := make([]int, len(line))
		copy(t, line)
		if simpleSafe(remove(t, i)) {
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
