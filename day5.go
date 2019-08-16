package main

//A nice string is one with all of the following properties:
// It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
// It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("day5.txt")

	reader := bufio.NewReader(file)

	niceStrings := 0
	niceStrings2 := 0

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

		if condition1(line) && condition2(line) && condition3(line) {
			niceStrings++
		}

		if condition4(line) && condition5(line) {
			niceStrings2++
		}
		fmt.Println("line : ", line)
	}

	fmt.Println("Number of Nice Strings : ", niceStrings)
	fmt.Println("Number of Nice Strings - Part 2 : ", niceStrings2)
}

// Part 1 Conditions
func condition1(line string) bool {
	cond := 0
	vowels := "aeiou"

	for _, i := range line {
		if strings.Contains(vowels, string(i)) {
			cond++
			if cond >= 3 {
				return true
			}
		}
	}

	return false
}

func condition2(line string) bool {
	for i, c := range line {
		if i < len(line)-1 && string(c) == string(line[i+1]) {
			return true
		}
	}

	return false
}

func condition3(line string) bool {
	unpermitted := []string{"ab", "cd", "pq", "xy"}

	for _, pattern := range unpermitted {
		if strings.Contains(line, pattern) {
			return false
		}
	}

	return true
}

// Part 2 Conditions

func condition4(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			return true
		}
	}

	return false
}

func condition5(str string) bool {
	for i := 0; i < len(str)-2; i++ {
		if strings.Count(str, str[i:i+2]) >= 2 {
			return true
		}
	}

	return false
}
