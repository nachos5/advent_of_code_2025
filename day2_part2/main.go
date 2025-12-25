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

			all_equal := false

			// loopa í gegnum string length og tékka á öllum mögulegum jöfnum skiptingum
			for i := range curr_string_len {
				if i == 0 {
					// repeated at least twice...
					continue
				}
				// ef enginn afgangur þá er hægt að skipta í jafna parta
				can_divide := curr_string_len%(i+1) == 0
				if !can_divide {
					continue
				}
				split_len := curr_string_len / (i + 1)
				first_split := curr_string[:split_len]
				// compare all other parts to first one
				all_equal_local := true
				for j := range i {
					next_split := curr_string[split_len*(j+1) : split_len*(j+2)]
					// fmt.Println(first_split, next_split)
					if first_split != next_split {
						all_equal_local = false
						break
					}
					// fmt.Println("first next", first_split, next_split)
				}
				// hægt að skipta upp í jafna parta allavega einu sinni
				if all_equal_local {
					all_equal = true
					break
				}
			}

			if all_equal {
				// fmt.Println("all equal", curr)
				total_sum += curr
			}

			curr++
		}
	}

	fmt.Println("Total sum", total_sum)

	elapsed := time.Since(start)
	fmt.Println("Time", elapsed)
}
