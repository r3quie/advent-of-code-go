package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

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
	slices.Sort(one)
	slices.Sort(two)

	var res int
	for i := 0; i < len(one); i++ {
		x := one[i] - two[i]
		if x < 0 {
			x = -x
		}
		res += x
	}
	fmt.Println(res)
}
