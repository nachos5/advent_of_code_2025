package main

import (
	"advent_of_code_2025/utils"
	"fmt"
	"strconv"
	_ "strings"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("day3 main")

	lines, err := utils.ReadFileLines("day3/input.txt")
	if err != nil {
		panic(err)
	}

	totalSum := 0

	for _, line := range lines {
		// for _, line := range lines[:1] {
		intList := lineToIntList(line)
		intListLen := len(intList)
		// skoðum ekki síðustu töluna
		maxValFirst, maxValIndexFirst := maxAndIndex(intList[:intListLen-1])
		maxValSecond, _ := maxAndIndex(intList[maxValIndexFirst+1:])
		// fmt.Println(intList[:intListLen-1])
		// fmt.Println(intList[maxValIndexFirst+1:])
		// fmt.Println(maxValFirst, maxValIndexFirst, maxValSecond)
		finalValue, err := strconv.Atoi(string(maxValFirst+'0') + string(maxValSecond+'0'))
		if err != nil {
			panic(err)
		}
		// fmt.Println(finalValue)
		totalSum += finalValue
	}

	fmt.Println("total sum", totalSum)

	elapsed := time.Since(start)
	fmt.Println("Time", elapsed)
}

func lineToIntList(line string) []int {
	int_list := []int{}
	for _, char := range line {
		int_list = append(int_list, int(char-'0'))
	}

	return int_list
}

func maxAndIndex(ints []int) (int, int) {
	curMax := 0
	curMaxIndex := 0

	for i, val := range ints {
		if val > curMax {
			curMax = val
			curMaxIndex = i
		}
	}

	return curMax, curMaxIndex

}
