package main

import (
	"advent_of_code_2025/utils"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("main start")

	lines, err := utils.ReadFileLines("day1/input.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Println(lines)
	var value = 50
	var total_zeroes = 0
	for _, line := range lines {
		// left - should add
		var should_add = line[0] == 'R'
		// string to int
		line_value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		// fmt.Println(should_add, value)
		if should_add {
			value += line_value
		} else {
			value -= line_value
		}
		// fmt.Println(value)
		if value%100 == 0 {
			total_zeroes++
		}
	}

	fmt.Println(value, total_zeroes)

	fmt.Println("main end")
}
