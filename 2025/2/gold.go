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
	data := strings.Split(string(dat), ",")
	foundNum :=0
	for _, x := range data {
		var numbersInt []int
		numbersStr := strings.Split(x, "-")

		tempVar, _  := strconv.Atoi(numbersStr[0])
		numbersInt = append(numbersInt, tempVar)
		tempVar, _ = strconv.Atoi(numbersStr[1])
		numbersInt = append(numbersInt, tempVar)
		
		counter := numbersInt[1] - numbersInt[0]

		for counter > -1 {
			tempNum := strconv.Itoa(numbersInt[1]-counter)
			strLen := len(tempNum)

			
			if(tempNum[strLen/2:] == tempNum[:strLen/2]) {
				y, _ := strconv.Atoi(tempNum)
				foundNum += y
				counter--
				continue
			}

			iterator := strLen
			for iterator > 0 {
				countAppearances := strings.Count(tempNum, tempNum[0:iterator])

				if(countAppearances*len(tempNum[0:iterator]) == strLen &&  countAppearances > 1) {
					y, _ := strconv.Atoi(tempNum)
					foundNum += y
					break
				}
				iterator--
			}
			counter--
		}

	}
	fmt.Println(foundNum)
}
