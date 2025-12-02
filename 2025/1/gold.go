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

func main() {
	dat, err := os.ReadFile("input.txt")
	check(err)
	data := strings.Split(string(dat), "\n")

	dialPos := 50
	result := 0

	for _, x := range data {
		if(len(x) < 2) {
			continue
		}

		number, _ := strconv.Atoi(x[1:])
		
		if(x[0] == 'L') {
			for number > 0 {
				dialPos--
				number--
				if(dialPos %100 == 0 || dialPos == 0) {
					result++
				}
			}
		}
		if(x[0] == 'R') {
			for number > 0 {
				dialPos++
				number--
				if(dialPos %100 == 0 || dialPos == 0) {
					result++
				}
			}
		}
	}

	fmt.Println(result)
}
