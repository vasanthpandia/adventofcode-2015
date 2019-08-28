package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var lights [1000][1000]bool

func main() {
	file, _ := os.Open("day6.txt")
	instrbytes, _ := ioutil.ReadAll(file)
	instructions := strings.Split(string(instrbytes), "\n")
	fmt.Println("Instruction Length: ", len(instructions))

	for i := 0; i <= 300; i++ {
		fetchInstructions(instructions[i])
	}

	lit := 0

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if lights[i][j] {
				lit++
			}
		}
	}

	fmt.Println("Lights lit : ", lit)
}

func processInstruction(num int, from [2]int, to [2]int) {
	for i := from[0]; i <= to[0]; i++ {
		for j := from[0]; j <= to[1]; j++ {
			if num == -1 {
				lights[i][j] = !lights[i][j]
			} else if num == 0 {
				lights[i][j] = false
			} else {
				lights[i][j] = true
			}
		}
	}
}

func fetchInstructions(instr string) {
	instrmap := make(map[string]int)
	instrmap["toggle"] = -1
	instrmap["on"] = 1
	instrmap["off"] = 0

	iwords := strings.Split(instr, " ")
	if len(iwords) == 5 {
		from := buildCoordinates(iwords[2])
		to := buildCoordinates(iwords[4])
		if iwords[1] == "on" {
			processInstruction(1, from, to)
		} else if iwords[1] == "off" {
			processInstruction(0, from, to)
		}
	} else if len(iwords) == 4 {
		from := buildCoordinates(iwords[1])
		to := buildCoordinates(iwords[3])
		processInstruction(-1, from, to)
	}
}

func buildCoordinates(word string) [2]int {
	coord := strings.Split(word, ",")
	x, err := strconv.Atoi(coord[0])
	if err != nil { fmt.Println(x, err) }
	out := strings.Fields(coord[1])
	y, err := strconv.Atoi(out[0])
	if err != nil { fmt.Println(y, err) }

	fmt.Println(x, y)

	return [2]int{x, y}
}
