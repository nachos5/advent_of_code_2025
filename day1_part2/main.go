package main

import (
	"advent_of_code_2025/utils"
	"fmt"
	"math"
	"strconv"
)

func main() {
	// test()
	// return

	fmt.Println("main start")

	lines, err := utils.ReadFileLines("day2/input.txt")
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

		if line_value == 0 {
			continue
		}

		// loop value before change
		var value_temp = value

		// fmt.Println(should_add, value)
		if should_add {
			value += line_value
		} else {
			value -= line_value
		}

		var value_temp_divides = value_temp / 100
		var value_divides = value / 100

		// ekki sama - hversu oft fór yfir 0 er mismunurinn.
		// dæmi: 456 l200 = 256 | 456 / 100 = 4, 256 / 100 = 2 | 4 - 2 = 2
		if value_temp_divides != value_divides {
			total_zeroes += int(math.Abs(float64(value_temp_divides) - float64(value_divides)))
		}

		// byrjar á 0 og mínusa - draga einn frá
		// dæmi: 400 l101 = 299 | 400 / 100 = 4, 299 / 100 = 2 | 4 - 2 = 2 en fer bara einu sinni yfir - mínusa 1
		if value_temp%100 == 0 && !should_add {
			total_zeroes -= 1
		}

		// endar á 0 - bæta við
		// dæmi: 400 l100 = 300 | 400 / 100 = 4, 300 / 100 = 3 | 4 - 3 = 1 - er rétt en reglan fyrir ofan tekur 1 frá svo bæta við .
		if value%100 == 0 && !should_add {
			total_zeroes += 1
		}

		// fmt.Println(total_zeroes)

		// fmt.Println(value_temp, value, value_temp - value, value_temp / 100, value / 100)

	}

	fmt.Println(value, total_zeroes)

	fmt.Println("main end")
}

func test() {
	var temp = 456 / 100
	fmt.Println(temp)
	fmt.Println(400 % 100)
}
