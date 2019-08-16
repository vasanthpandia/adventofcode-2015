package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	floorNumber()
}

func floorNumber() {
	f, _ := os.Open("day1.txt")
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	currentFloor := 0
	var charPosition int
	fmt.Println("CharPosition : ", charPosition)

	for i := 0; i < len(content); i++ {
		if content[i] == 40 {
			currentFloor++
		} else if content[i] == 41 {
			currentFloor--
		}

		if currentFloor == -1 && charPosition == 0 {
			charPosition = i + 1
		}
	}

	fmt.Println("Final floor is : ", currentFloor)
	fmt.Println("CharPosition : ", charPosition)
}
