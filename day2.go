package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"io"
	"bufio"
	"sort"
)

func main() {
	sheetArea()
}

func sheetArea() {
	file, _ := os.Open("day2.txt")

	reader := bufio.NewReader(file)

	var finalArea float64

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}

		if len(line) <= 1 {
			break
		}
		fmt.Println("line : ", line)

		dimensions := strings.Split(line, "x")

		side1, _ := strconv.ParseFloat(strings.TrimSpace(dimensions[0]), 64)
		side2, _ := strconv.ParseFloat(strings.TrimSpace(dimensions[1]), 64)
		side3, _ := strconv.ParseFloat(strings.TrimSpace(dimensions[2]), 64)

		fmt.Println(side1, side2, side3)

		finalArea += findArea(side1, side2, side3)
	}

	fmt.Printf("Final Area : %f", finalArea)
}
