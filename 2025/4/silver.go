package main

import (
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkForAdjecent(strArray []string, x int, y int) bool {
	counter := 0
	xPlus := false
	xMinus := false
	yPlus := false
	yMinus := false

	if x+1 < len(strArray[y]) {
		xPlus = true
	}
	if x-1 >= 0 {
		xMinus = true
	}
	if y-1 >= 0 {
		yMinus = true
	}
	if y+1 < len(strArray) {
		yPlus = true
	}

	if xPlus {
		if strArray[y][x+1] == '@' {
			counter++
		}
		if yPlus {
			if strArray[y+1][x+1] == '@' {
				counter++
			}
		}
		if yMinus {
			if strArray[y-1][x+1] == '@' {
				counter++
			}
		}
	}

	if xMinus {
		if strArray[y][x-1] == '@' {
			counter++
		}
		if yPlus {
			if strArray[y+1][x-1] == '@' {
				counter++
			}
		}
		if yMinus {
			if strArray[y-1][x-1] == '@' {
				counter++
			}
		}
	}

	if yPlus {
		if strArray[y+1][x] == '@' {
			counter++
		}
	}

	if yMinus {
		if strArray[y-1][x] == '@' {
			counter++
		}
	}

	if counter < 4 {
		return true
	}
	return false

}

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	data := strings.Split(string(dat), "\n")
	counter := 0
	//data[y][x]

	for y, _ := range data {
		for x, _ := range data[y] {
			if data[y][x] == '@' {
				if checkForAdjecent(data, x, y) {
					counter++
				}
			}
		}
	}
	println(counter)
}
