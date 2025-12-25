package main

import (
	"advent_of_code_2025/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("day2 main")

	lines_temp, err := utils.ReadFileLines("day2/input.txt")
	if err != nil {
		panic(err)
	}
	var lines = strings.Split(lines_temp[0], ",")
	// fmt.Println(lines)
	// fmt.Println(len(lines))

	total_sum := 0

	for _, line := range lines {
		// for _, line := range lines[:1] {
		line_split := strings.Split(line, "-")
		start, err := strconv.Atoi(line_split[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(line_split[1])
		if err != nil {
			panic(err)
		}
		if start >= end {
			fmt.Println("start >= end")
			continue
		}

		curr := start
		for curr <= end {
			curr_string := strconv.Itoa(curr)
			curr_string_len := len(curr_string)
			// ef oddatala þá er ekki hægt að splitta í tvo jafna strengi
			if curr_string_len%2 == 1 {
				curr++
				continue
			}
			// splitta í tvo jafna parta
			var first_part = curr_string[:curr_string_len/2]
			var second_part = curr_string[curr_string_len/2 : curr_string_len]
			// bæta við ef fyrri parturinn er eins og seinni
			if first_part == second_part {
				total_sum += curr
			}
			// fmt.Println(first_part, second_part)
			curr++
		}
	}

	fmt.Println("Total sum", total_sum)

	elapsed := time.Since(start)
	fmt.Println("Time", elapsed)
}
