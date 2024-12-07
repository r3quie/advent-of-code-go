package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumOfElements(arr []int) map[int]int {
	d := make(map[int]int)
	for _, num := range arr {
		d[num] = d[num] + 1
	}
	return d
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var one, two []int // would probably be more efficient to use make() here; depends if you can get len() without reading the whole file
	lines := bufio.NewScanner(file)
	for lines.Scan() {
		s := strings.Split(lines.Text(), "   ")
		x, _ := strconv.Atoi(s[0])
		y, _ := strconv.Atoi(s[1])

		one = append(one, x)
		two = append(two, y)
	}
	var res int

	d := getNumOfElements(two)
	for _, x := range one {
		res += x * d[x]
	}
	fmt.Println(res)
}
