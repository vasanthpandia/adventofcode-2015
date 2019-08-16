package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sheetArea()
}

func sheetArea() {
	file, _ := os.Open("day2.txt")

	reader := bufio.NewReader(file)

	var finalArea float64
	var finalRibbonLength float64

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
		finalRibbonLength += ribbonLength(side1, side2, side3)
	}

	fmt.Printf("Final Area : %f\n", finalArea)
	fmt.Printf("Final Ribbon Length : %f\n", finalRibbonLength)
}

func findArea(side1 float64, side2 float64, side3 float64) float64 {
	area1 := side1 * side2
	area2 := side2 * side3
	area3 := side1 * side3

	area := 2 * (area1 + area2 + area3)

	var smallest float64

	if area1 <= area2 && area1 <= area3 {
		smallest = area1
	}
	if area2 <= area1 && area2 <= area3 {
		smallest = area2
	}
	if area3 <= area1 && area3 <= area2 {
		smallest = area3
	}

	return (area + smallest)
}

func ribbonLength(side1 float64, side2 float64, side3 float64) float64 {
	dims := []float64{side1, side2, side3}
	sort.Float64s(dims)

	smallPeri := 2 * (dims[0] + dims[1])

	return smallPeri + volume(side1, side2, side3)
}

func volume(side1 float64, side2 float64, side3 float64) float64 {
	return side1 * side2 * side3
}
