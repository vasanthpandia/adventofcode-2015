package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	findHouseCount()
}

func findHouseCount() {
	file, _ := os.Open("day3.txt")
	reader := bufio.NewReader(file)
	content, _ := ioutil.ReadAll(reader)

	houseSetYear1 := NewIntPointSet()
	HouseCount(content, &houseSetYear1)
	fmt.Println("Total for Year 1 : ", houseSetYear1.NumElements())

	santa, robosanta := SplitInputs(content)
	houseSetYear2 := NewIntPointSet()
	HouseCount(santa, &houseSetYear2)
	HouseCount(robosanta, &houseSetYear2)
	fmt.Println("Total for Year 2: ", houseSetYear2.NumElements())
}

func HouseCount(content []byte, houseSet *IntPointSet) {
	x, y := 0, 0

	for _, val := range content {
		if string(val) == ">" {
			x++
			s := IntPoint{
				X: x,
				Y: y,
			}
			houseSet.Add(s)
		} else if string(val) == "<" {
			x--
			s := IntPoint{
				X: x,
				Y: y,
			}
			houseSet.Add(s)
		} else if string(val) == "^" {
			y++
			s := IntPoint{
				X: x,
				Y: y,
			}
			houseSet.Add(s)
		} else if string(val) == "v" {
			y--
			s := IntPoint{
				X: x,
				Y: y,
			}
			houseSet.Add(s)
		}
	}
}

// Split the given instructions between Santa and RoboSanta

func SplitInputs(content []byte) ([]byte, []byte) {
	var santa []byte
	var robosanta []byte

	for i, val := range content {
		if i%2 == 0 {
			santa = append(santa, val)
		} else {
			robosanta = append(robosanta, val)
		}
	}

	return santa, robosanta
}

// Set Implementation - Thanks to @amahfouz - https://softwareengineering.stackexchange.com/a/273372

// types

type IntPoint struct {
	X, Y int
}

// set implementation for small number of items
type IntPointSet struct {
	slice []IntPoint
}

// functions

func (p1 IntPoint) Equals(p2 IntPoint) bool {
	return (p1.X == p2.X) && (p1.Y == p2.Y)
}

func (set *IntPointSet) Add(p IntPoint) {
	if !set.Contains(p) {
		set.slice = append(set.slice, p)
	}
}

func (set IntPointSet) Contains(p IntPoint) bool {
	for _, v := range set.slice {
		if v.Equals(p) {
			return true
		}
	}
	return false
}

func (set IntPointSet) NumElements() int {
	return len(set.slice)
}

func NewIntPointSet() IntPointSet {
	return IntPointSet{(make([]IntPoint, 0, 10))}
}
