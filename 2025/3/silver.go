package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convStrtoIntArray(str string) []int {
	var array []int

	for _, x := range str {
		y, _ :=strconv.Atoi(string(x))
		array = append(array, y)
	}

	return array
}

func findHigh(array []int, length int)(int, int){
	highestIndex :=0
	highestNum :=0
	for x, _ := range array {
		if(array[x] > highestNum && len(array)-length > x) {
			highestNum = array[x]
			highestIndex = x
		}
	}
	fmt.Println(highestNum)
	return highestIndex, highestNum
}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	data := strings.Split(string(dat), "\n")
	total :=0

	for _, x := range data {
		array := convStrtoIntArray(x)
		highIndex, number := findHigh(array, 1)
		number *=10
		highIndex, num2 := findHigh(array[highIndex+1:], 0)
		number += num2
		total += number
		fmt.Println("total ", total)
	}

	fmt.Println(total)


}
